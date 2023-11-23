\c OCPP;

-- Создаем таблицу ChargerPoints
CREATE TABLE IF NOT EXISTS ChargerPoints (
     id SERIAL PRIMARY KEY,
     web_id VARCHAR(100) NOT NULL,
     name VARCHAR(255) NOT NULL,
     last_heartbeat TIMESTAMP,
     location VARCHAR(255) NOT NULL,
     ocpp_version VARCHAR(20) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     is_deleted BOOLEAN
);

-- Создаем таблицу Connectors
CREATE TABLE IF NOT EXISTS Connectors (
      id SERIAL PRIMARY KEY,
      chargerpoint_id INT NOT NULL,
      connector_number INT NOT NULL,
      status VARCHAR(20) NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      is_deleted BOOLEAN,
      FOREIGN KEY (chargerpoint_id) REFERENCES ChargerPoints(id) ON DELETE CASCADE
);