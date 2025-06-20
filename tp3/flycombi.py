#!/usr/bin/python3

import sys
import func_aux

def main():
    ruta_aeropuertos = sys.argv[1]
    ruta_vuelos = sys.argv[2]
    aeropuertos_por_ciudad, aeropuertos_por_codigo = func_aux.leer_aeropuertos(ruta_aeropuertos)
    lista_vuelos = func_aux.leer_vuelos(ruta_vuelos)
    grafo_aeropuertos = func_aux.crear_grafo_aeropuertos(aeropuertos_por_codigo, lista_vuelos)
    ultimo_camino = None
    for entrada in sys.stdin:
        lista_entrada = entrada.strip().split(" ", 1)
        nuevo_ultimo_camino = func_aux.realizar_accion_segun_comando(lista_entrada, ultimo_camino, aeropuertos_por_ciudad, aeropuertos_por_codigo, grafo_aeropuertos)
        if nuevo_ultimo_camino is not None:
            ultimo_camino = nuevo_ultimo_camino

if __name__ == "__main__":
    main()