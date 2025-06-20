from grafo import Grafo
from aeropuerto import Aeropuerto
from vuelos import Vuelo
import biblioteca

def leer_aeropuertos(ruta):
    aeropuertos_por_ciudad, aeropuertos_por_codigo = {}, {}
    with open(ruta) as f:
        for linea in f:
            linea_leida = linea.split(",")
            ciudad, codigo_aerop, latitud, longitud = linea_leida[0], linea_leida[1], linea_leida[2], linea_leida[3]
            aeropuerto = Aeropuerto(ciudad, codigo_aerop, latitud, longitud)
            aeropuertos_por_ciudad[ciudad] = aeropuertos_por_ciudad.get(ciudad, [])
            aeropuertos_por_ciudad[ciudad].append(aeropuerto)
            aeropuertos_por_codigo[codigo_aerop] = aeropuerto
    return aeropuertos_por_ciudad, aeropuertos_por_codigo

def leer_vuelos(ruta):
    lista_vuelos = []
    with open(ruta) as f:
        for linea in f:
            linea_leida = linea.split(",")
            vuelo = Vuelo(linea_leida[0], linea_leida[1], linea_leida[2], linea_leida[3], linea_leida[4])
            lista_vuelos.append(vuelo)
    return lista_vuelos

def crear_grafo_aeropuertos(aeropuertos_por_codigo, lista_vuelos):
    grafo = Grafo(False)
    for codigo_aerop in aeropuertos_por_codigo.keys():
        grafo.agregar_vertice(codigo_aerop)
    for vuelo in lista_vuelos:
        aerop_i, aerop_j = vuelo.ver_aeropuerto_origen(), vuelo.ver_aeropuerto_destino()
        if not grafo.estan_unidos(aerop_i, aerop_j):
            grafo.agregar_arista(aerop_i, aerop_j, vuelo)
    return grafo

def crear_grafo_ciudades(archivo_ciudades):
    grafo = Grafo(True)
    with open(archivo_ciudades) as f:
        lineas = f.readlines()
        ciudades = lineas[0].rstrip().split(",")
        for ciudad in ciudades:
            grafo.agregar_vertice(ciudad)
        for i in range(1, len(lineas)):
            ciudad_1, ciudad_2 = lineas[i].rstrip().split(",")
            grafo.agregar_arista(ciudad_1, ciudad_2, 1)
    return grafo

def sumar_sin_infinitos(distancia):
    suma = 0
    for distancia in distancia.values():
        if distancia != float("inf"):
            suma += distancia
    return suma

def elegir_mejor_camino(aeropuertos_por_ciudad, grafo_aeropuertos, criterio, ciudad_origen, ciudad_destino):
    aeropuertos_en_origen, aeropuertos_en_destino = aeropuertos_por_ciudad[ciudad_origen], aeropuertos_por_ciudad[ciudad_destino]
    mejor_camino_padre = None
    for aerop_origen in aeropuertos_en_origen:
        for aerop_destino in aeropuertos_en_destino:
            if criterio == "barato":
                padre, distancia = biblioteca.camino_mas(grafo_aeropuertos, aerop_origen.ver_codigo(), aerop_destino.ver_codigo(), lambda vuelo: int(vuelo.ver_precio()))
            elif criterio == "rapido":
                padre, distancia = biblioteca.camino_mas(grafo_aeropuertos, aerop_origen.ver_codigo(), aerop_destino.ver_codigo(), lambda vuelo: int(vuelo.ver_tiempo_promedio()))
            else:
                padre, distancia = biblioteca.camino_minimo_escalas(grafo_aeropuertos, aerop_origen.ver_codigo(), aerop_destino.ver_codigo())
            if mejor_camino_padre is None or sumar_sin_infinitos(distancia) < mejor_suma:
                mejor_camino_padre, mejor_destino = padre, aerop_destino.ver_codigo()
                mejor_suma = sumar_sin_infinitos(distancia)
    return mejor_camino_padre, mejor_destino

def reconstruir_camino(padre, destino):
    camino = []
    actual = destino
    while actual is not None:
        camino.append(actual)
        actual = padre[actual]
    camino.reverse()
    return camino

def ejecutar_camino_mas(lista_entrada, aeropuertos_por_ciudad, grafo_aeropuertos):
    _, parametros = lista_entrada
    lista_parametros = parametros.split(",")
    criterio, ciudad_origen, ciudad_destino = lista_parametros
    padre, aerop_destino = elegir_mejor_camino(aeropuertos_por_ciudad, grafo_aeropuertos, criterio, ciudad_origen, ciudad_destino)
    camino = reconstruir_camino(padre, aerop_destino)
    print(" -> ".join(camino))
    return camino

def ejecutar_camino_escalas(lista_entrada, aeropuertos_por_ciudad, grafo_aeropuertos):
    _, parametros = lista_entrada
    lista_parametros = parametros.split(",")
    ciudad_origen, ciudad_destino = lista_parametros
    padre, aerop_destino = elegir_mejor_camino(aeropuertos_por_ciudad, grafo_aeropuertos, "escalas", ciudad_origen, ciudad_destino)
    camino = reconstruir_camino(padre, aerop_destino)
    print(" -> ".join(camino))
    return camino

def ejecutar_centralidad(lista_entrada, grafo_aeropuertos):
    _, n = lista_entrada
    vertices_mas_centrales = biblioteca.centralidad(grafo_aeropuertos)
    n_vertices_mas_centrales = vertices_mas_centrales[:int(n)]
    print(", ".join(n_vertices_mas_centrales))

def escribir_archivo_nueva_aerolinea(ruta, mst):
    with open(ruta, "w") as f:
        for arista in mst.obtener_aristas():
            aerop_origen, aerop_destino, vuelo = arista
            tiempo_prom, precio, cant_vuelos_entre_aeropuertos = vuelo.ver_tiempo_promedio(), vuelo.ver_precio(), vuelo.ver_cant_vuelos_entre_aeropuertos()
            f.write(f"{aerop_origen},{aerop_destino},{tiempo_prom},{precio},{cant_vuelos_entre_aeropuertos}")
    
def ejecutar_nueva_aerolinea(lista_entrada, grafo_aeropuertos):
    _, ruta_archivo = lista_entrada
    mst = biblioteca.mst_prim(grafo_aeropuertos)
    escribir_archivo_nueva_aerolinea(ruta_archivo, mst)
    print("OK")

def ejecutar_itinerario(lista_entrada, grafo_aeropuertos, aeropuertos_por_ciudad):
    _, ruta_archivo = lista_entrada
    grafo_ciudades = crear_grafo_ciudades(ruta_archivo)
    ciudades_a_visitar = biblioteca.orden_topologico(grafo_ciudades)
    print(", ".join(ciudades_a_visitar))
    for i in range(len(ciudades_a_visitar)):
        if i+1 == len(ciudades_a_visitar):
            break
        padre, destino = elegir_mejor_camino(aeropuertos_por_ciudad, grafo_aeropuertos, "rapido", ciudades_a_visitar[i], ciudades_a_visitar[i+1])
        camino = reconstruir_camino(padre, destino)
        print(" -> ".join(camino))

def guardar_coordenadas(ultimo_camino, aeropuertos_por_codigo):
    coords = {}
    for cod_aeropuerto in ultimo_camino:
        latitud, longitud = aeropuertos_por_codigo[cod_aeropuerto].ver_latitud(), aeropuertos_por_codigo[cod_aeropuerto].ver_longitud()
        coords[cod_aeropuerto] = (latitud, longitud)
    return coords

def ejecutar_exportar_kml(lista_entrada, ultimo_camino, aeropuertos_por_codigo):
    _, ruta_archivo = lista_entrada
    coords = guardar_coordenadas(ultimo_camino, aeropuertos_por_codigo)
    encabezado_kml = '''<?xml version='1.0' encoding='UTF-8'?>\n'''
    inicio_kml = '''<kml xmlns='http://earth.google.com/kml/2.1'>
      <Document>\n'''
    info_kml = '''        <name>Ultimo Mejor Camino</name>
        <description>Se muestra el ultimo mejor camino encontrado.</description>\n'''
    contenido_kml = encabezado_kml + inicio_kml + info_kml
    for i, cod_aerop in enumerate(ultimo_camino):
        coords_act = f"<coordinates>{coords[cod_aerop][1]}".rstrip("\n") + f", {coords[cod_aerop][0]}</coordinates>".rstrip("\n")
        contenido_kml += f'''
        <Placemark>
            <name>{cod_aerop}</name>
            <Point>
                {coords_act}
            </Point>
        </Placemark>\n'''
    for i in range(len(ultimo_camino)):
        if i+1 == len(ultimo_camino):
            break
        coords_act = f"<coordinates>{coords[ultimo_camino[i]][1]}".rstrip("\n") + f", {coords[ultimo_camino[i]][0]} {coords[ultimo_camino[i+1]][1]}".rstrip("\n") + f", {coords[ultimo_camino[i+1]][0]}</coordinates>".rstrip("\n")
        contenido_kml += f'''
        <Placemark>
            <LineString>
                {coords_act}
            </LineString>
        </Placemark>\n'''
    fin_kml = '''\n      </Document>\n</kml>'''
    contenido_kml += fin_kml
    with open(ruta_archivo, 'w') as f:
        f.write(contenido_kml)
    print("OK")

def realizar_accion_segun_comando(lista_entrada, ultimo_camino, aeropuertos_por_ciudad, aeropuertos_por_codigo, grafo_aeropuertos):
    comando = lista_entrada[0]
    nuevo_ultimo_camino = None
    if comando == "camino_mas":
        nuevo_ultimo_camino = ejecutar_camino_mas(lista_entrada, aeropuertos_por_ciudad, grafo_aeropuertos)
    if comando == "camino_escalas":
        nuevo_ultimo_camino = ejecutar_camino_escalas(lista_entrada, aeropuertos_por_ciudad, grafo_aeropuertos)
    if comando == "centralidad":
        ejecutar_centralidad(lista_entrada, grafo_aeropuertos)
    if comando == "nueva_aerolinea":
        ejecutar_nueva_aerolinea(lista_entrada, grafo_aeropuertos)
    if comando == "itinerario":
        ejecutar_itinerario(lista_entrada, grafo_aeropuertos, aeropuertos_por_ciudad)
    if comando == "exportar_kml":
        ejecutar_exportar_kml(lista_entrada, ultimo_camino, aeropuertos_por_codigo)
    return nuevo_ultimo_camino