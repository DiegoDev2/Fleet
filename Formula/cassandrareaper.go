package main

// CassandraReaper - Management interface for Cassandra
// Homepage: https://cassandra-reaper.io/

import (
	"fmt"
	
	"os/exec"
)

func installCassandraReaper() {
	// Método 1: Descargar y extraer .tar.gz
	cassandrareaper_tar_url := "https://github.com/thelastpickle/cassandra-reaper/releases/download/3.6.1/cassandra-reaper-3.6.1-release.tar.gz"
	cassandrareaper_cmd_tar := exec.Command("curl", "-L", cassandrareaper_tar_url, "-o", "package.tar.gz")
	err := cassandrareaper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cassandrareaper_zip_url := "https://github.com/thelastpickle/cassandra-reaper/releases/download/3.6.1/cassandra-reaper-3.6.1-release.zip"
	cassandrareaper_cmd_zip := exec.Command("curl", "-L", cassandrareaper_zip_url, "-o", "package.zip")
	err = cassandrareaper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cassandrareaper_bin_url := "https://github.com/thelastpickle/cassandra-reaper/releases/download/3.6.1/cassandra-reaper-3.6.1-release.bin"
	cassandrareaper_cmd_bin := exec.Command("curl", "-L", cassandrareaper_bin_url, "-o", "binary.bin")
	err = cassandrareaper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cassandrareaper_src_url := "https://github.com/thelastpickle/cassandra-reaper/releases/download/3.6.1/cassandra-reaper-3.6.1-release.src.tar.gz"
	cassandrareaper_cmd_src := exec.Command("curl", "-L", cassandrareaper_src_url, "-o", "source.tar.gz")
	err = cassandrareaper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cassandrareaper_cmd_direct := exec.Command("./binary")
	err = cassandrareaper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
