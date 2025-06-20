class Vuelo:
    def __init__(self, aerop_origen, aerop_destino, tiempo_promedio, precio, cant_vuelos_entre_aeropuertos):
        self.aerop_origen = aerop_origen
        self.aerop_destino = aerop_destino
        self.tiempo_promedio = tiempo_promedio
        self.precio = precio
        self.cant_vuelos_entre_aeropuertos = cant_vuelos_entre_aeropuertos
    
    def ver_aeropuerto_origen(self):
        return self.aerop_origen
    
    def ver_aeropuerto_destino(self):
        return self.aerop_destino

    def ver_tiempo_promedio(self):
        return self.tiempo_promedio

    def ver_precio(self):
        return self.precio
    
    def ver_cant_vuelos_entre_aeropuertos(self):
        return self.cant_vuelos_entre_aeropuertos

    def __lt__(self, otra):
        return int(self.precio) < int(otra.precio)