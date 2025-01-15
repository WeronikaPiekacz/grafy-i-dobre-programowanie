# Eulerian Graph Solver

Program służący do znajdowania cyklu Eulera w grafie nieskierowanym.

## Funkcjonalności

-   Tworzenie grafu na podstawie listy wierzchołków i krawędzi
-   Znajdowanie cyklu Eulera (jeśli istnieje)
-   Konfiguracja poprzez plik config.json
-   Walidacja grafu pod kątem parzystości stopni wierzchołków (w przygotowaniu)

## Użycie

1. Przygotuj plik konfiguracyjny config.json
2. Zdefiniuj graf poprzez listę wierzchołków i krawędzi
3. Uruchom program aby znaleźć cykl Eulera

## Przykład

```go
graph := bl.NewGraph([]string{"0", "1", "2"}, []bl.Edge{
    *bl.NewEdge("0", "1"),
    *bl.NewEdge("1", "2"),
    *bl.NewEdge("2", "0"),
})

result, err := graph.FindCircuit()
```

## Konfiguracja

Plik config.json pozwala na:

-   Możliwość odczytu grafu z pliku tekstowego

## Wymagania

-   Go 1.20 lub nowszy
