# Guia de Desenvolvimento - Frontend Angular

Este documento descreve a estrutura de arquivos e o processo para criar novas telas (screens) e serviços no projeto Angular.

## Estrutura do Projeto

```
src/app/
├── components/          # Componentes reutilizáveis
├── guards/             # Guards de rota (autenticação, autorização)
├── layout/             # Componentes de layout (header, sidebar, etc.)
├── models/             # Interfaces e modelos TypeScript
├── screens/            # Telas/páginas da aplicação
├── services/           # Serviços (API, utilitários)
├── app.config.ts       # Configuração da aplicação
├── app.routes.ts       # Definição de rotas
├── app.html            # Template principal
├── app.scss            # Estilos globais
└── app.ts              # Componente principal
```

## Criando uma Nova Tela (Screen)

### 1. Gerar Componente com Angular CLI

Use o Angular CLI para gerar automaticamente a estrutura da tela:

```bash
# Gerar nova tela/screen
ng generate component screens/nome-da-tela.screen --skip-tests

```

Isso criará automaticamente:
```
src/app/screens/nome-da-tela/
├── nome-da-tela.screen.ts      # Componente TypeScript
├── nome-da-tela.screen.html    # Template HTML
└── nome-da-tela.screen.scss    # Estilos específicos
```

E atualize o selector e templateUrl no arquivo `.ts`

### 2. Template do Componente TypeScript (.ts)

```typescript
import { Component, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
// Importe os serviços necessários
import { SeuService } from '../../services/seu.service';

@Component({
  selector: 'app-nome-da-tela',
  templateUrl: './nome-da-tela.screen.html',
  styleUrls: ['./nome-da-tela.screen.scss'],
  imports: [
    CommonModule,
    ReactiveFormsModule
    // Adicione outros módulos conforme necessário
  ]
})
export class NomeDaTelaScreen {
  // Propriedades do componente
  isLoading = signal(false);
  errorMessage: string = '';
  
  // Formulários (se necessário)
  form: FormGroup;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private seuService: SeuService
  ) {
    // Inicialização do formulário (se necessário)
    this.form = this.fb.group({
      campo: ['', [Validators.required]]
    });
  }

  // Métodos do componente
  onSubmit(): void {
    if (this.form.invalid) {
      return;
    }
    
    this.isLoading.set(true);
    // Lógica do submit
    console.log('Submit realizado com sucesso!');
  }
}
```

### 3. Template HTML (.html)

```html
<div class="container">
  <div class="header">
    <h1>Título da Tela</h1>
  </div>

  <div class="content">
    <!-- Conteúdo da tela -->
    <form [formGroup]="form" (ngSubmit)="onSubmit()" *ngIf="form">
      <!-- Campos do formulário -->
      <div class="form-group">
        <label for="campo">Campo:</label>
        <input 
          type="text" 
          id="campo" 
          formControlName="campo"
          class="form-control"
        >
      </div>

      <!-- Mensagem de erro -->
      <div class="error-message" *ngIf="errorMessage">
        {{ errorMessage }}
      </div>

      <!-- Botão de submit -->
      <button 
        type="submit" 
        class="btn btn-primary"
        [disabled]="isLoading() || form.invalid"
      >
        <span *ngIf="isLoading()">Carregando...</span>
        <span *ngIf="!isLoading()">Enviar</span>
      </button>
    </form>
  </div>
</div>
```

### 4. Estilos SCSS (.scss)

```scss
.container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;

  .header {
    margin-bottom: 30px;
    
    h1 {
      color: #333;
      font-size: 2rem;
    }
  }

  .content {
    .form-group {
      margin-bottom: 20px;

      label {
        display: block;
        margin-bottom: 5px;
        font-weight: 500;
      }

      .form-control {
        width: 100%;
        padding: 10px;
        border: 1px solid #ddd;
        border-radius: 4px;
        font-size: 14px;

        &:focus {
          outline: none;
          border-color: #007bff;
          box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
        }
      }
    }

    .error-message {
      color: #dc3545;
      font-size: 14px;
      margin-bottom: 15px;
    }

    .btn {
      padding: 10px 20px;
      border: none;
      border-radius: 4px;
      cursor: pointer;
      font-size: 14px;

      &.btn-primary {
        background-color: #007bff;
        color: white;

        &:hover:not(:disabled) {
          background-color: #0056b3;
        }

        &:disabled {
          background-color: #6c757d;
          cursor: not-allowed;
        }
      }
    }
  }
}
```

### 5. Adicionar Rota

No arquivo `src/app/app.routes.ts`, adicione a nova rota:

```typescript
import { Routes } from '@angular/router';
import { NomeDaTelaScreen } from './screens/nome-da-tela/nome-da-tela.screen';

export const routes: Routes = [
  // ... outras rotas
  {
    path: 'nome-da-rota',
    component: NomeDaTelaScreen
  }
];
```

## Criando um Novo Serviço

### 1. Gerar Serviço com Angular CLI

Use o Angular CLI para gerar automaticamente o serviço:

```bash
# Gerar novo serviço
ng generate service services/nome-do-servico --skip-tests
```

Isso criará automaticamente o arquivo `src/app/services/nome-do-servico.service.ts`

### 2. Template do Serviço

```typescript
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject, throwError } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
// Importe os modelos necessários
import { SeuModel } from '../models/seu-model.model';

@Injectable({
  providedIn: 'root'
})
export class NomeDoServicoService {
  private apiUrl = '/api'; // URL base da API
  
  // Subject para estado reativo (opcional)
  private dataSubject = new BehaviorSubject<SeuModel[]>([]);
  public data$ = this.dataSubject.asObservable();

  constructor(private http: HttpClient) {}

  // Método GET
  getAll(): Observable<SeuModel[]> {
    return this.http.get<SeuModel[]>(`${this.apiUrl}/endpoint`).pipe(
      tap(data => this.dataSubject.next(data)),
      catchError(this.handleError)
    );
  }

  // Método GET por ID
  getById(id: number): Observable<SeuModel> {
    return this.http.get<SeuModel>(`${this.apiUrl}/endpoint/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  // Método POST
  create(data: Partial<SeuModel>): Observable<SeuModel> {
    return this.http.post<SeuModel>(`${this.apiUrl}/endpoint`, data).pipe(
      tap(newItem => {
        const currentData = this.dataSubject.value;
        this.dataSubject.next([...currentData, newItem]);
      }),
      catchError(this.handleError)
    );
  }

  // Método PUT
  update(id: number, data: Partial<SeuModel>): Observable<SeuModel> {
    return this.http.put<SeuModel>(`${this.apiUrl}/endpoint/${id}`, data).pipe(
      tap(updatedItem => {
        const currentData = this.dataSubject.value;
        const index = currentData.findIndex(item => item.id === id);
        if (index !== -1) {
          currentData[index] = updatedItem;
          this.dataSubject.next([...currentData]);
        }
      }),
      catchError(this.handleError)
    );
  }

  // Método DELETE
  delete(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/endpoint/${id}`).pipe(
      tap(() => {
        const currentData = this.dataSubject.value;
        const filteredData = currentData.filter(item => item.id !== id);
        this.dataSubject.next(filteredData);
      }),
      catchError(this.handleError)
    );
  }

  // Tratamento de erros
  private handleError(error: any): Observable<never> {
    console.error('Erro no serviço:', error);
    let errorMessage = 'Erro interno do servidor';
    
    if (error.error?.message) {
      errorMessage = error.error.message;
    } else if (error.message) {
      errorMessage = error.message;
    }
    
    return throwError(() => new Error(errorMessage));
  }
}
```

## Modelos (Models)

### Criando um Modelo

Crie o arquivo em `src/app/models/nome-do-modelo.model.ts`:

```typescript
export interface NomeDoModelo {
  id: number;
  campo1: string;
  campo2: number;
  campo3?: boolean; // Campo opcional
  createdAt: Date;
  updatedAt: Date;
}


```

## Checklist para Nova Funcionalidade

### Para uma Nova Tela:
- [ ] Executar `ng g c screens/nome-da-tela.screen --skip-tests`
- [ ] Customizar o componente TypeScript
- [ ] Customizar o template HTML
- [ ] Customizar os estilos SCSS
- [ ] Adicionar rota em `app.routes.ts`
- [ ] Testar navegação e funcionalidade

### Para um Novo Serviço:
- [ ] Executar `ng g s services/nome-do-servico --skip-tests`
- [ ] Implementar métodos necessários (GET, POST, PUT, DELETE)
- [ ] Adicionar tratamento de erros
- [ ] Criar interfaces/modelos relacionados
- [ ] Testar integração com componentes

### Boas Práticas:
- [ ] Usar nomenclatura consistente (kebab-case para arquivos)
- [ ] Implementar tratamento de erros adequado
- [ ] Usar signals para estado reativo
- [ ] Adicionar validação de formulários
- [ ] Implementar loading states
- [ ] Seguir padrões de design estabelecidos
- [ ] Documentar código complexo

## Comandos Úteis

```bash
# Gerar novo componente
ng generate component screens/nome-da-tela --skip-tests

# Gerar novo serviço
ng generate service services/nome-do-servico --skip-tests

# Executar aplicação
ng serve

# Build para produção
ng build --prod
```

## Configuração do Proxy

O arquivo `src/proxy.conf.json` configura o proxy para a API durante o desenvolvimento:

```json
{
  "/api/*": {
    "target": "http://localhost:3000",
    "secure": true,
    "changeOrigin": true
  }
}
```
