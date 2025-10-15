# Riviera Beauty Interiors - Yacht Provisioning Management System

This is an internal management system for **Riviera Beauty Interiors**, a specialized provisioning company serving the luxury yacht industry.

## Business Context

Riviera Beauty Interiors manages orders for yacht owners, handling everything from luxury linens and cosmetics to personalized amenities. The company acts as a single point of contact, consolidating orders from multiple suppliers into unified deliveries for yacht owners.

## Project Vision

This application is being developed as part of a Master's program to modernize and streamline the entire business workflow. The system will evolve in phases:

### Phase 1: Inventory Management (In Development)

A comprehensive internal inventory management system that helps track stock, organize items, monitor supplier details, and prevent shortages.

### Phase 2: Task Management (Planned)

A customized, user-friendly task management system (mini-Jira) tailored specifically for the team's workflow. Unlike generic project management tools, this will be:

- Designed with non-technical users in mind
- Integrated directly with inventory data
- Focused on order fulfillment workflows
- Connected to active orders and delivery schedules

### Phase 3: Client Portal (Future)

A dedicated portal for yacht owners to:

- Track delivery status in real-time
- View order history
- Place quick reorders of previously purchased items
- Streamline communication between clients and the provisioning team

---

## Planned Features (Phase 1)

The inventory management system will provide:

- **Folders** to group items by product categories, brands, or suppliers
- **Items** with comprehensive tracking:
  - Quantity, weight, volume
  - Expiry dates and shelf life
  - Supplier pricing and order details
  - Minimum stock levels with automated alerts
  - Custom notes and specifications
- **Suppliers** database with full contact information for quick reordering
- **Usages** tracking to record stock changes (sales, usage, returns)
- **Photos** for visual documentation and identification
- **Audit logs** for accountability and change tracking
- **Statistics and Alerts**:
  - Low stock notifications (items below minimum threshold)
  - Upcoming expiry warnings
  - Usage trends and shortage history
  - Most frequently ordered items

---

## Tech Stack

### Frontend

- **TanStack Start** with **React** for the user interface
- Modern, responsive design for both desktop and mobile use

### Backend

- **Go** with **Gin** framework for API server
- **PostgreSQL** database for data persistence

---

## Getting Started

### Prerequisites

- Go 1.21+ for backend development
- Node.js 20+ and pnpm/npm for frontend development
- PostgreSQL 15+ for database

### Development Setup

_(Development instructions will be added as the project structure is finalized)_

---

## Core Concepts

### Inventory Management

- **Folders**: Organize items by categories (e.g., "Bathroom Amenities", "Linens", "Cosmetics")
- **Items**: Individual products with full tracking capabilities
- **Suppliers**: Contact database for procurement
- **Minimum Level**: Automated low-stock alerts when quantity drops below threshold
- **Usages**: Historical record of all stock movements
- **Audit Logs**: Complete accountability trail for all system changes

### Business Workflow

1. Yacht owner places order
2. Staff checks inventory availability
3. Items below minimum level trigger supplier reorder
4. Usage records track fulfillment
5. Statistics inform purchasing decisions

### Planned Features (Phase 2 & 3)

- Order management integrated with task tracking
- Delivery scheduling and status updates
- Client portal for order tracking and reordering
- Automated notification system for clients
- Enhanced reporting for business analytics

---

## Development Roadmap

### 🔄 Phase 1: Inventory Management (In Development)

- Core inventory tracking
- Supplier management
- Low stock alerts
- Usage history
- Basic statistics

### 📋 Phase 2: Task Management (Planned)

- Custom task board (simplified Jira alternative)
- Order-to-task linking
- Team collaboration features
- Delivery scheduling
- Integrated with inventory system

### 🚀 Phase 3: Client Portal (Future)

- Client authentication and profiles
- Real-time delivery tracking
- Order history viewing
- Quick reorder functionality
- Client-staff messaging

---

## Project Context

This application is being developed as a Master's degree project, designed to modernize operations for a real business with specific needs:

- Non-technical user base requiring intuitive interfaces
- Yacht industry specifics (luxury goods, time-sensitive deliveries)
- Multi-stakeholder system (internal staff + external clients)
- Scalability for future business growth

---

## License

This application is proprietary to **Riviera Beauty Interiors**.  
Unauthorized copying or distribution is prohibited.
