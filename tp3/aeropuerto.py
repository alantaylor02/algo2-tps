class Aeropuerto:
    def __init__(self, ciudad, codigo, latitud, longitud):
        self.ciudad = ciudad
        self.codigo = codigo
        self.latitud = latitud
        self.longitud = longitud
    
    def ver_ciudad(self):
        return self.ciudad
    
    def ver_codigo(self):
        return self.codigo
    
    def ver_latitud(self):
        return self.latitud
    
    def ver_longitud(self):
        return self.longitud