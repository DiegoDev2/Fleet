package main

// SchemaEvolutionManager - Manage postgresql database schema migrations
// Homepage: https://github.com/mbryzek/schema-evolution-manager

import (
	"fmt"
	
	"os/exec"
)

func installSchemaEvolutionManager() {
	// Método 1: Descargar y extraer .tar.gz
	schemaevolutionmanager_tar_url := "https://github.com/mbryzek/schema-evolution-manager/archive/refs/tags/0.9.54.tar.gz"
	schemaevolutionmanager_cmd_tar := exec.Command("curl", "-L", schemaevolutionmanager_tar_url, "-o", "package.tar.gz")
	err := schemaevolutionmanager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	schemaevolutionmanager_zip_url := "https://github.com/mbryzek/schema-evolution-manager/archive/refs/tags/0.9.54.zip"
	schemaevolutionmanager_cmd_zip := exec.Command("curl", "-L", schemaevolutionmanager_zip_url, "-o", "package.zip")
	err = schemaevolutionmanager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	schemaevolutionmanager_bin_url := "https://github.com/mbryzek/schema-evolution-manager/archive/refs/tags/0.9.54.bin"
	schemaevolutionmanager_cmd_bin := exec.Command("curl", "-L", schemaevolutionmanager_bin_url, "-o", "binary.bin")
	err = schemaevolutionmanager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	schemaevolutionmanager_src_url := "https://github.com/mbryzek/schema-evolution-manager/archive/refs/tags/0.9.54.src.tar.gz"
	schemaevolutionmanager_cmd_src := exec.Command("curl", "-L", schemaevolutionmanager_src_url, "-o", "source.tar.gz")
	err = schemaevolutionmanager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	schemaevolutionmanager_cmd_direct := exec.Command("./binary")
	err = schemaevolutionmanager_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
