package main

// Ivy - Agile dependency manager
// Homepage: https://ant.apache.org/ivy/

import (
	"fmt"
	
	"os/exec"
)

func installIvy() {
	// Método 1: Descargar y extraer .tar.gz
	ivy_tar_url := "https://www.apache.org/dyn/closer.lua?path=ant/ivy/2.5.2/apache-ivy-2.5.2-bin.tar.gz"
	ivy_cmd_tar := exec.Command("curl", "-L", ivy_tar_url, "-o", "package.tar.gz")
	err := ivy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ivy_zip_url := "https://www.apache.org/dyn/closer.lua?path=ant/ivy/2.5.2/apache-ivy-2.5.2-bin.zip"
	ivy_cmd_zip := exec.Command("curl", "-L", ivy_zip_url, "-o", "package.zip")
	err = ivy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ivy_bin_url := "https://www.apache.org/dyn/closer.lua?path=ant/ivy/2.5.2/apache-ivy-2.5.2-bin.bin"
	ivy_cmd_bin := exec.Command("curl", "-L", ivy_bin_url, "-o", "binary.bin")
	err = ivy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ivy_src_url := "https://www.apache.org/dyn/closer.lua?path=ant/ivy/2.5.2/apache-ivy-2.5.2-bin.src.tar.gz"
	ivy_cmd_src := exec.Command("curl", "-L", ivy_src_url, "-o", "source.tar.gz")
	err = ivy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ivy_cmd_direct := exec.Command("./binary")
	err = ivy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
