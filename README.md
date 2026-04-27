# 🚀 Micro Frontend Demo with React + Angular

This project demonstrates a **Micro Frontend Architecture** using **Webpack Module Federation**. It consists of 4 independent applications (3 React + 1 Angular) that work together:

## 📁 Project Structure

```
react-angular/
├── host/                    # Container/Host Application - React (Port 3000)
├── product-list/           # Product List Micro Frontend - React (Port 3001)
├── product-form/           # Product Form Micro Frontend - React (Port 3002)
├── product-stats/          # Product Stats Micro Frontend - Angular (Port 3003)
└── README.md
```

## 🎯 What is Micro Frontend?

Micro Frontends extend the concept of microservices to frontend development. Each part of the UI is:
- **Independent** - Can be developed, tested, and deployed separately
- **Technology Agnostic** - Mix React, Angular, Vue in the same app! ⚡
- **Scalable** - Teams work autonomously without stepping on each other

## 🏗️ Architecture

```
┌──────────────────────────────────────────────────────┐
│              HOST APP (3000) - React                 │
│  ┌────────────────────────────────────────────────┐  │
│  │      Module Federation Container               │  │
│  └────────────────────────────────────────────────┘  │
│         ↓              ↓              ↓               │
│   ┌─────────┐    ┌─────────┐    ┌─────────┐        │
│   │ Product │    │ Product │    │ Product │        │
│   │  Form   │    │  List   │    │  Stats  │        │
│   │ (React) │    │ (React) │    │(Angular)│        │
│   │  :3002  │    │  :3001  │    │  :3003  │        │
│   └─────────┘    └─────────┘    └─────────┘        │
└──────────────────────────────────────────────────────┘
```
│    ┌─────────────┐    ┌─────────────┐      │
│    │ ProductForm │    │ ProductList │      │
│    │   (3002)    │    │   (3001)    │      │
│    └─────────────┘    └─────────────┘      │
└─────────────────────────────────────────────┘
```

## 🚀 Getting Started

### Prerequisites
- Node.js (v14 or higher)
- npm or yarn

### Installation

**Option 1: Install all at once (Recommended)**

```bash
# From react-angular directory
npm install
```

**Option 2: Install individually**

```bash
# Install Host
cd host
npm install

# Install Product List
cd ../product-list
npm install

# Install Product Form
cd ../product-form
npm install

# Install Product Stats (Angular)
cd ../product-stats
npm install
```

## ▶️ Running the Application

You need to run **all 4 applications** simultaneously:

### Terminal 1 - Product List (Port 3001)
```bash
cd product-list
npm start
```

### Terminal 2 - Product Form (Port 3002)
```bash
cd product-form
npm start
```

### Terminal 3 - Product Stats Angular (Port 3003)
```bash
cd product-stats
npm start
```

### Terminal 4 - Host Application (Port 3000)
```bash
cd host
npm start
```

### 🌐 Access the Applications

- **Host Application (Main)**: http://localhost:3000
- **Product List (Standalone)**: http://localhost:3001
- **Product Form (Standalone)**: http://localhost:3002
- **Product Stats Angular (Standalone)**: http://localhost:3003

## ✨ Features

### Product Form Micro Frontend (React)
- Add new products with name and price
- Form validation
- Responsive design
- Can run standalone or integrated

### Product List Micro Frontend (React)
- Display all products in a beautiful grid
- Shows total product count
- Empty state handling
- Can run standalone or integrated

### Product Stats Micro Frontend (Angular) 🅰️
- Real-time statistics calculation
- Total products, value, and average price
- Most expensive and cheapest product tracking
- Beautiful gradient cards
- Demonstrates React + Angular integration

### Host Application (React)
- Integrates all micro frontends (React + Angular)
- Shared state management
- Lazy loading with Suspense
- Beautiful gradient UI

## 🔧 Technology Stack

- **React 18** - UI Library (Host, Product Form, Product List)
- **Angular 17** - Framework (Product Stats)
- **Webpack 5** - Module bundler
- **Module Federation** - Micro Frontend integration
- **Babel** - JavaScript compiler (React apps)
- **TypeScript** - Type safety (Angular app)
- **CSS3** - Styling

## 📚 How Module Federation Works

### 1. **Product Form exposes its component:**
```javascript
// webpack.config.js in product-form
exposes: {
  './ProductForm': './src/ProductForm',
}
```

### 2. **Product List exposes its component:**
```javascript
// webpack.config.js in product-list
exposes: {
  './ProductList': './src/ProductList',
}
```

### 3. **Host consumes both:**
```javascript
// webpack.config.js in host
remotes: {
  productList: 'productList@http://localhost:3001/remoteEntry.js',
  productForm: 'productForm@http://localhost:3002/remoteEntry.js',
}
```

### 4. **Host imports and uses them:**
```javascript
const ProductList = lazy(() => import('productList/ProductList'));
const ProductForm = lazy(() => import('productForm/ProductForm'));
```

## 🎨 Key Benefits Demonstrated

1. **Independent Development** - Each MFE can be developed separately
2. **Independent Deployment** - Deploy form without affecting list
3. **Shared Dependencies** - React is shared (singleton) to avoid duplication
4. **Lazy Loading** - Components loaded only when needed
5. **Technology Freedom** - Could easily swap React for Vue/Angular in one MFE

## 🛠️ Development

### Run in Standalone Mode

Each micro frontend can run independently for development:

```bash
# Test Product Form alone
cd product-form
npm start
# Visit http://localhost:3002

# Test Product List alone
cd product-list
npm start
# Visit http://localhost:3001
```

### Build for Production

```bash
# Build all apps
cd host && npm run build
cd ../product-list && npm run build
cd ../product-form && npm run build
```

## 🐛 Troubleshooting

### Issue: "Failed to load remote entry"
**Solution**: Make sure all 3 apps are running before accessing the host

### Issue: Port already in use
**Solution**: Kill the process using that port or change port in webpack.config.js

### Issue: Module not found
**Solution**: Run `npm install` in each directory

## 📖 Learning Resources

- [Module Federation Docs](https://webpack.js.org/concepts/module-federation/)
- [Micro Frontend Architecture](https://martinfowler.com/articles/micro-frontends.html)

## 🎯 Next Steps

Try these exercises to learn more:
1. Add a delete product feature
2. Add product categories
3. Create a third MFE for product search
4. Add state management (Redux/Context)
5. Deploy each MFE to separate domains

---

**Happy Coding! 🚀**
