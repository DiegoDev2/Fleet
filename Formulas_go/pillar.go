package main

// Pillar - Manage migrations for Cassandra data stores
// Homepage: https://github.com/comeara/pillar

import (
	"fmt"
	
	"os/exec"
)

func installPillar() {
	// Método 1: Descargar y extraer .tar.gz
	pillar_tar_url := "https://github.com/comeara/pillar/archive/refs/tags/v2.3.0.tar.gz"
	pillar_cmd_tar := exec.Command("curl", "-L", pillar_tar_url, "-o", "package.tar.gz")
	err := pillar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pillar_zip_url := "https://github.com/comeara/pillar/archive/refs/tags/v2.3.0.zip"
	pillar_cmd_zip := exec.Command("curl", "-L", pillar_zip_url, "-o", "package.zip")
	err = pillar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pillar_bin_url := "https://github.com/comeara/pillar/archive/refs/tags/v2.3.0.bin"
	pillar_cmd_bin := exec.Command("curl", "-L", pillar_bin_url, "-o", "binary.bin")
	err = pillar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pillar_src_url := "https://github.com/comeara/pillar/archive/refs/tags/v2.3.0.src.tar.gz"
	pillar_cmd_src := exec.Command("curl", "-L", pillar_src_url, "-o", "source.tar.gz")
	err = pillar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pillar_cmd_direct := exec.Command("./binary")
	err = pillar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbt")
exec.Command("latte", "install", "sbt")
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
}
