import random

class Grafo:
    def __init__(self, dirigido):
        self.grafo = {}
        self.dirigido = dirigido

    def agregar_vertice(self, v):
        if v not in self.grafo:
            self.grafo[v] = {}

    def borrar_vertice(self, v):
        if v in self.grafo:
            self.grafo.pop(v)
            for w in self.grafo:
                if v in self.grafo[w]:
                    self.grafo[w].pop(v)

    def agregar_arista(self, v, w, peso=1):
        if v in self.grafo and w in self.grafo:
            self.grafo[v][w] = peso
            if not self.dirigido:
                self.grafo[w][v] = peso

    def borrar_arista(self, v, w):
        if v in self.grafo and w in self.grafo[v]:
            self.grafo[v].pop(w)
            if not self.dirigido:
                self.grafo[w].pop(v)

    def estan_unidos(self, v, w):
        return v in self.grafo and w in self.grafo[v]

    def peso_arista(self, v, w):
        if self.estan_unidos(v, w):
            return self.grafo[v][w]
        return None
    
    def obtener_vertices(self):
        return list(self.grafo.keys())
    
    def obtener_aristas(self):
        aristas = []
        for v in self.grafo:
            for w in self.adyacentes(v):
                aristas.append((v, w, self.peso_arista(v, w)))
        return aristas
    
    def vertice_aleatorio(self):
        return random.choice(self.obtener_vertices())
    
    def adyacentes(self, v):
        return list(self.grafo[v].keys())