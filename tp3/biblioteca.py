from grafo import Grafo
from collections import deque
import heapq

def camino_mas(grafo, origen, destino, func_criterio):
    distancia, padre = {}, {}
    for v in grafo.obtener_vertices():
        distancia[v] = float("inf")
    distancia[origen] = 0
    padre[origen] = None
    heap = []
    heapq.heappush(heap, (0, origen))
    while len(heap) != 0:
        _, v = heapq.heappop(heap)
        if v == destino:
            return padre, distancia
        for w in grafo.adyacentes(v):
            vuelo = grafo.peso_arista(v, w)
            dato = func_criterio(vuelo)
            if distancia[v] + dato < distancia[w]:
                distancia[w] = distancia[v] + dato
                padre[w] = v
                heapq.heappush(heap, (distancia[w], w))
    return padre, distancia

def camino_minimo_escalas(grafo, origen, destino):
    distancia, padre, visitado = {}, {}, {}
    distancia[origen] = 0
    padre[origen] = None
    visitado[origen] = True
    q = deque()
    q.appendleft(origen)
    while len(q) != 0:
        v = q.pop()
        if v == destino:
            return padre, distancia
        for w in grafo.adyacentes(v):
            if w not in visitado:
                distancia[w] = distancia[v] + 1
                padre[w] = v
                visitado[w] = True
                q.appendleft(w)
    return padre, distancia

def camino_maximo(grafo, origen, func_criterio):
    distancia, padre = {}, {}
    for v in grafo.obtener_vertices():
        distancia[v] = float("inf")
    distancia[origen] = 0
    padre[origen] = None
    heap = []
    heapq.heappush(heap, (0, origen))
    while len(heap) != 0:
        _, v = heapq.heappop(heap)
        for w in grafo.adyacentes(v):
            vuelo = grafo.peso_arista(v, w)
            dato_inv = func_criterio(vuelo)
            if distancia[v] + dato_inv < distancia[w]:
                distancia[w] = distancia[v] + dato_inv
                padre[w] = v
                heapq.heappush(heap, (distancia[w], w))
    return padre, distancia

def ordenar_vertices(distancias):
    vertices_distancias = list(distancias.items())
    vertices_distancias.sort(key=lambda x: x[1], reverse=True)
    vertices_ordenados = []
    for v, distancia in vertices_distancias:
        if distancia != float("inf"):
            vertices_ordenados.append(v)
    return vertices_ordenados
    
def centralidad(grafo):
    cent = {}
    for v in grafo.obtener_vertices(): 
        cent[v] = 0
    for v in grafo.obtener_vertices():
        padre, distancia = camino_maximo(grafo, v, lambda vuelo: 1 / int(vuelo.ver_cant_vuelos_entre_aeropuertos()))
        cent_aux = {}
        for w in grafo.obtener_vertices(): 
            cent_aux[w] = 0
        vertices_ordenados = ordenar_vertices(distancia)
        for w in vertices_ordenados:
            if not padre[w] is None:
                cent_aux[padre[w]] += 1 + cent_aux[w]
        for w in grafo.obtener_vertices():
            if w == v: 
                continue
            cent[w] += cent_aux[w]
    cent_ordenado = sorted(cent, key=cent.get, reverse=True)
    return cent_ordenado

def mst_prim(grafo):
    v_inicial = grafo.vertice_aleatorio()
    visitados = set()
    visitados.add(v_inicial)
    heap = []
    mst = Grafo(False)
    for v in grafo.obtener_vertices():
        mst.agregar_vertice(v)
    for w in grafo.adyacentes(v_inicial):
        vuelo = grafo.peso_arista(v_inicial, w)
        heapq.heappush(heap, (vuelo, (v_inicial, w)))
    while len(heap) != 0:
        vuelo, (u, w) = heapq.heappop(heap)
        if w not in visitados:
            mst.agregar_arista(u, w, vuelo)
            visitados.add(w)
            for x in grafo.adyacentes(w):
                if x not in visitados:
                    vuelo = grafo.peso_arista(w, x)
                    heapq.heappush(heap, (vuelo, (w, x)))
    return mst

def pila_a_lista(pila):
    res = []
    while len(pila) != 0:
        res.append(pila.pop())
    return res

def _orden_topologico(grafo, v, pila, visitados):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            _orden_topologico(grafo, w, pila, visitados)
    pila.append(v)

def orden_topologico(grafo):
    visitados = set()
    pila = deque()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            _orden_topologico(grafo, v, pila, visitados)
    return pila_a_lista(pila)