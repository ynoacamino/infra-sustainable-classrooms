#!/bin/bash

# Script para generar el código del microservicio de autenticación con Goa

echo "🚀 Generando código del microservicio de autenticación con TOTP..."

# Navegar al directorio del servicio de auth
cd /home/ynoacamino/dev/infrastructure/services/auth

# Generar el código usando goa gen
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/auth/design/api
