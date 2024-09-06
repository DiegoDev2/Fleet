package main

// Mongroup - Monitor a group of processes with mon
// Homepage: https://github.com/jgallen23/mongroup

import (
	"fmt"
	
	"os/exec"
)

func installMongroup() {
	// Método 1: Descargar y extraer .tar.gz
	mongroup_tar_url := "https://github.com/jgallen23/mongroup/archive/refs/tags/0.4.1.tar.gz"
	mongroup_cmd_tar := exec.Command("curl", "-L", mongroup_tar_url, "-o", "package.tar.gz")
	err := mongroup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongroup_zip_url := "https://github.com/jgallen23/mongroup/archive/refs/tags/0.4.1.zip"
	mongroup_cmd_zip := exec.Command("curl", "-L", mongroup_zip_url, "-o", "package.zip")
	err = mongroup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongroup_bin_url := "https://github.com/jgallen23/mongroup/archive/refs/tags/0.4.1.bin"
	mongroup_cmd_bin := exec.Command("curl", "-L", mongroup_bin_url, "-o", "binary.bin")
	err = mongroup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongroup_src_url := "https://github.com/jgallen23/mongroup/archive/refs/tags/0.4.1.src.tar.gz"
	mongroup_cmd_src := exec.Command("curl", "-L", mongroup_src_url, "-o", "source.tar.gz")
	err = mongroup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongroup_cmd_direct := exec.Command("./binary")
	err = mongroup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mon")
exec.Command("latte", "install", "mon")
}
