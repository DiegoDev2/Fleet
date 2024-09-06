package main

// Ant - Java build tool
// Homepage: https://ant.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installAnt() {
	// Método 1: Descargar y extraer .tar.gz
	ant_tar_url := "https://www.apache.org/dyn/closer.lua?path=ant/binaries/apache-ant-1.10.15-bin.tar.xz"
	ant_cmd_tar := exec.Command("curl", "-L", ant_tar_url, "-o", "package.tar.gz")
	err := ant_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ant_zip_url := "https://www.apache.org/dyn/closer.lua?path=ant/binaries/apache-ant-1.10.15-bin.tar.xz"
	ant_cmd_zip := exec.Command("curl", "-L", ant_zip_url, "-o", "package.zip")
	err = ant_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ant_bin_url := "https://www.apache.org/dyn/closer.lua?path=ant/binaries/apache-ant-1.10.15-bin.tar.xz"
	ant_cmd_bin := exec.Command("curl", "-L", ant_bin_url, "-o", "binary.bin")
	err = ant_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ant_src_url := "https://www.apache.org/dyn/closer.lua?path=ant/binaries/apache-ant-1.10.15-bin.tar.xz"
	ant_cmd_src := exec.Command("curl", "-L", ant_src_url, "-o", "source.tar.gz")
	err = ant_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ant_cmd_direct := exec.Command("./binary")
	err = ant_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
