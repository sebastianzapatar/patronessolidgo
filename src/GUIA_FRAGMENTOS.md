# Guía para escribir fragmentos de slides (patrones de diseño en Go)

La presentación final `index.html` se arma concatenando los archivos de `src/` en orden
alfabético (`00_head.html`, `10_...`, `20_...`, ..., `99_foot.html`). Cada fragmento contiene
**solo** bloques `<div class="slide">...</div>` (sin `<html>`, `<head>`, `<style>` ni `<script>`).
Todos los estilos ya existen en `src/00_head.html` — léelo y usa SOLO esas clases
(se permite `style="..."` inline pequeño cuando haga falta).

Idioma: **español** (código Go con identificadores en español cuando sea natural, comentarios en español).

## Atributos obligatorios de cada slide

```html
<div class="slide" data-section="Creacionales · Singleton">
```

- `data-section`: texto que se muestra en la barra inferior.
- En la **primera** slide de cada patrón añade además `id="p-<slug>"` y `data-toc="<Nombre>"`
  y `data-toc-group="Creacionales|Estructurales|Comportamiento"` (esto genera el índice lateral).
- Cada slide empieza con `<span class="slide-number-badge"></span>` (se numera solo por JS).

## Estructura por patrón: 3 slides

### Slide 1 — Teoría (qué problema resuelve)
```html
<div class="slide" id="p-factory-method" data-toc="Factory Method" data-toc-group="Creacionales" data-section="Creacionales · Factory Method">
  <span class="slide-number-badge"></span>
  <div class="slide-inner">
    <div class="pattern-head">
      <span class="cat-badge creacional">🏭 Creacional</span>
      <span class="part-badge">1 / 3 · Teoría</span>
    </div>
    <h2 class="slide-title">Factory Method</h2>
    <p class="slide-subtitle">También llamado: Método fábrica, Constructor virtual</p>

    <div class="intent-box">
      <span class="label">🎯 Propósito</span>
      <p>...</p>
    </div>

    <div class="two-cols">
      <div class="glass-card problem-card"><h3>😟 Problema</h3><p>...</p></div>
      <div class="glass-card solution-card"><h3>💡 Solución</h3><p>...</p></div>
    </div>

    <div class="analogy">
      <div class="icon">🌍</div>
      <div><h3>Analogía en el mundo real</h3><p>...</p></div>
    </div>
  </div>
</div>
```

### Slide 2 — Estructura, aplicabilidad, pros/contras
- `.uml` con un diagrama ASCII de la estructura (cajas con `┌─┐ │ └─┘`, flechas `──▶`, `◁──` para
  implementación, etc.). Escapa `<` y `>` como `&lt;` `&gt;`.
- `<ol class="participants">` explicando cada participante numerado del diagrama.
- `glass-card` "🧭 Aplicabilidad" (cuándo usarlo) y "🛠️ Cómo implementarlo" (pasos).
- `.pros-cons` con `.pros` / `.cons`.
- Opcional: "🔗 Relación con otros patrones".
- Badge: `<span class="part-badge">2 / 3 · Estructura</span>`.

### Slide 3 — Implementación en Go
- `two-cols-wide`: a la izquierda el código completo, a la derecha la explicación paso a paso
  (`ol`) + la salida real + un `.go-tip` con una nota idiomática de Go.
- Código:
  ```html
  <pre class="code-block" data-lang="go" data-label="factorymethod/main.go">package main
  ...
  }</pre>
  ```
  - El código va justo después de `>` y `</pre>` va pegado a la última línea.
  - **ESCAPA** `&` → `&amp;`, `<` → `&lt;`, `>` → `&gt;` dentro del código (el resaltado de
    sintaxis se aplica automáticamente con JS; NO pongas `<span>` manuales).
- Salida: `<pre class="output-block">...salida exacta de go run...</pre>`.
- Enlace a la fuente al final:
  `<a class="source-link" href="https://refactoring.guru/es/design-patterns/factory-method" target="_blank">📖 Teoría basada en refactoring.guru — Factory Method</a>`
  (puede ir en la slide 1 o 3).
- Badge: `<span class="part-badge">3 / 3 · Go</span>`.

## Código Go

- Cada ejemplo es un programa ejecutable en `ejemplos/<categoria>/<slug>/main.go` (`package main`),
  dentro del módulo `patrones` (`ejemplos/go.mod`). Categorías: `creacionales`, `estructurales`,
  `comportamiento`.
- Verifícalo con `cd ejemplos && go vet ./<categoria>/<slug> && go run ./<categoria>/<slug>`.
- El código de la slide debe ser **idéntico** al archivo (o una versión con partes irrelevantes
  omitidas indicadas con `// ...`), y la salida mostrada debe ser la salida real.
- Ejemplos de 40–90 líneas: claros, didácticos, con comentarios que señalen cada participante
  del patrón (p. ej. `// Producto (interfaz)`, `// Creador concreto`).
- Usa Go idiomático: interfaces pequeñas, composición/embedding en lugar de herencia,
  constructores `NewX`, `sync.Once` para Singleton, funciones como valores cuando aplique, etc.
  Explica en el `.go-tip` cómo se adapta el patrón a Go (no hay clases ni herencia).

## Teoría

Basada en https://refactoring.guru/es/design-patterns/<slug> (lee la página). Parafrasea y explica
con tus palabras en español, de forma clara y didáctica (no copies párrafos literales largos).
Incluye propósito, problema, solución, analogía, estructura, aplicabilidad, pasos de implementación,
pros y contras y relaciones con otros patrones.
