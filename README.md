# 🐹 Patrones de Diseño y Principios SOLID en Go

Clase interactiva y repositorio de ejemplos ejecutables sobre **Principios SOLID** y los **22 Patrones de Diseño (GoF)** implementados de forma idiomática en el **lenguaje Go**, desarrollada para el curso de **Lenguajes y Compiladores (2026-2)**.

---

## ✨ Contenido del Proyecto

La presentación interactiva (`index.html`, **91 slides**) y el módulo de código (`ejemplos/`) cubren de principio a fin:

1. **Orientación a Objetos en Go sin Clases**:
   - Cómo Go reemplaza la herencia clásica por `struct`, métodos con receptor, **interfaces implícitas** (*structural typing*), composición por **embedding** y funciones de primera clase.
2. **Los 5 Principios SOLID en Go** *(basados en el caso de estudio de librería/facturas y estacionamiento de [freeCodeCamp](https://www.freecodecamp.org/espanol/news/los-principios-solid-explicados-en-espanol/))*:
   - **S** — *Single Responsibility Principle* (`ejemplos/solid/srp`)
   - **O** — *Open/Closed Principle* (`ejemplos/solid/ocp`)
   - **L** — *Liskov Substitution Principle* (`ejemplos/solid/lsp`)
   - **I** — *Interface Segregation Principle* (`ejemplos/solid/isp`)
   - **D** — *Dependency Inversion Principle* (`ejemplos/solid/dip`)
3. **Fundamentos de Patrones de Diseño e Ingeniería de Contexto con LLMs**:
   - Qué es un patrón vs. un algoritmo, historia (*Christopher Alexander* y *Gang of Four*), ventajas, críticas y clasificación.
   - **Patrones, buen orden del código y ahorro de tokens al usar LLMs**: cómo la modularidad (SRP), los contratos ligeros (ISP/DIP) y el vocabulario estándar de patrones reducen entre un **60 % y 80 %** el consumo de tokens en asistentes de IA.
4. **Catálogo completo de 22 Patrones de Diseño** *(basados en [Refactoring.Guru](https://refactoring.guru/es/design-patterns), con 3 slides por patrón: ① Teoría y analogía, ② Estructura UML, aplicabilidad y pros/contras, ③ Implementación ejecutable en Go)*:
   - 🏭 **Creacionales (5)**: *Factory Method, Abstract Factory, Builder, Prototype, Singleton*.
   - 🧱 **Estructurales (7)**: *Adapter, Bridge, Composite, Decorator, Facade, Flyweight, Proxy*.
   - 🧠 **Comportamiento (10)**: *Chain of Responsibility, Command, Iterator (`iter.Seq` Go 1.23+), Mediator, Memento, Observer, State, Strategy, Template Method, Visitor (AST de expresiones para compiladores)*.

---

## 🤖 ¿Por qué los Patrones de Diseño y el Buen Orden Ahorran Tokens en LLMs?

Cuando trabajamos con asistentes de código basados en IA (Claude, Gemini, Copilot, Cursor), **el desorden arquitectónico tiene un costo directo en tokens, latencia y precisión**:

1. **Compresión semántica del prompt (Vocabulario compartido)**:
   - Los nombres de los patrones funcionan como atajos de alta densidad semántica. Pedir *"aplica un Decorator sobre la interfaz `Notificador`"* comprime en una sola oración lo que de otro modo requeriría párrafos enteros de explicación y múltiples rondas de corrección.
2. **Aislamiento de contexto gracias a SRP y paquetes cohesivos**:
   - Un archivo monolítico (*God Object* de 1.500 líneas) obliga a inyectar miles de tokens irrelevantes en la ventana de contexto cada vez que pides un cambio. Al separar responsabilidades en archivos pequeños (40–90 líneas), el agente solo lee el archivo involucrado, reduciendo el consumo de tokens de entrada entre un **60 % y un 80 %**.
3. **Interfaces como contratos ligeros (DIP e ISP)**:
   - En Go, las interfaces pequeñas permiten pasarle al LLM únicamente el contrato (`type FacturaPersistencia interface { Guardar(Factura) error }`) de 3 líneas en lugar de las implementaciones completas de base de datos, caché o servicios externos.

### 📚 Referencias que lo respaldan
- **Anthropic Engineering (2025)** — [*Effective Context Engineering for AI Agents*](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents): explica cómo la arquitectura modular, archivos bien delimitados y la carga progresiva de contexto (*just-in-time context*) evitan contaminar la ventana de tokens y mejoran el rendimiento de los agentes de programación.
- **arXiv:2501.04835 (2025)** — [*Do Code LLMs Understand Design Patterns?*](https://arxiv.org/abs/2501.04835): demuestra empíricamente que los LLMs de código comprenden y aprovechan los patrones de diseño como abstracciones de alto nivel para generar soluciones más precisas y consistentes.
- **Liu et al., Stanford / TACL (2024)** — [*Lost in the Middle: How Language Models Use Long Contexts* (arXiv:2307.03172)](https://arxiv.org/abs/2307.03172): demuestra que saturar la ventana de contexto con código extenso o acoplado degrada la capacidad de atención del modelo en la mitad del contexto además de incrementar cuadráticamente el costo computacional.

---

## 🗂️ Estructura del Repositorio

```text
.
├── index.html                  # Presentación interactiva completa (lista para navegador / GitHub Pages)
├── src/                        # Fragmentos modulares HTML que componen index.html
│   ├── 00_head.html            # Sistema de diseño CSS (paleta lila / glassmorphism)
│   ├── 10_intro.html           # Portada, agenda y POO en Go sin clases
│   ├── 12_solid.html           # Los 5 principios SOLID explicados con código Go
│   ├── 15_patrones_intro.html  # Qué es un patrón, historia, críticas, clasificación y mapa
│   ├── 16_patrones_llm.html    # Patrones, orden del código y ahorro de tokens en LLMs
│   ├── 19_div_creacionales.html
│   ├── 20_creacionales.html    # 5 patrones creacionales (15 slides)
│   ├── 29_div_estructurales.html
│   ├── 30_estructurales.html   # 7 patrones estructurales (21 slides)
│   ├── 39_div_comportamiento.html
│   ├── 40_comportamiento_a.html# 5 patrones de comportamiento — Parte A (15 slides)
│   ├── 41_comportamiento_b.html# 5 patrones de comportamiento — Parte B (15 slides)
│   ├── 90_cierre.html          # Tabla rápida de decisión y cierre
│   ├── 99_foot.html            # Barra de navegación, índice lateral (M) y resaltador Go
│   └── GUIA_FRAGMENTOS.md      # Guía de estilo y estructura de las slides
└── ejemplos/                   # Módulo Go con los 27 programas ejecutables
    ├── go.mod
    ├── solid/                  # srp, ocp, lsp, isp, dip
    ├── creacionales/           # factorymethod, abstractfactory, builder, prototype, singleton
    ├── estructurales/          # adapter, bridge, composite, decorator, facade, flyweight, proxy
    └── comportamiento/         # chainofresponsibility, command, iterator, mediator, memento,
                                # observer, state, strategy, templatemethod, visitor
```

---

## 🚀 Cómo Usar la Presentación y Ejecutar los Ejemplos

### 1. Ver las diapositivas
Abre directamente `index.html` en cualquier navegador moderno:
- **`←` / `→` / `Espacio`**: avanzar o retroceder entre diapositivas.
- **Tecla `M`** (o botón `☰ Índice`): abrir el panel lateral con el índice interactivo de todos los temas y patrones.
- **`Inicio` / `Fin`**: ir a la primera o última diapositiva.

### 2. Ejecutar los ejemplos en Go
Todos los ejemplos están dentro del módulo `ejemplos/` y no requieren dependencias externas (solo la librería estándar de Go):

```bash
cd ejemplos

# Verificar que los 27 paquetes compilan sin advertencias
go vet ./...

# Ejecutar un principio SOLID (ej. Responsabilidad Única)
go run ./solid/srp

# Ejecutar un patrón creacional (ej. Factory Method)
go run ./creacionales/factorymethod

# Ejecutar un patrón estructural (ej. Decorator)
go run ./estructurales/decorator

# Ejecutar un patrón de comportamiento (ej. Visitor sobre un AST)
go run ./comportamiento/visitor
```

### 3. Reconstruir `index.html` tras editar `src/`
Si modificas algún fragmento dentro de `src/`, regenera el archivo principal concatenándolos en orden:

```bash
cat src/[0-9]*.html > index.html
```

---

## 📖 Créditos y Fuentes Teóricas
- Teoría y catálogo de patrones de diseño basados en [**Refactoring.Guru — Patrones de Diseño**](https://refactoring.guru/es/design-patterns) y [**Patrones en Go**](https://refactoring.guru/es/design-patterns/go).
- Explicación y caso de estudio de los principios SOLID basados en [**freeCodeCamp — Los principios SOLID explicados en español**](https://www.freecodecamp.org/espanol/news/los-principios-solid-explicados-en-espanol/).
