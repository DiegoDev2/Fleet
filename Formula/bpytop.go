package main

// Bpytop - Linux/OSX/FreeBSD resource monitor
// Homepage: https://github.com/aristocratos/bpytop

import (
	"fmt"
	
	"os/exec"
)

func installBpytop() {
	// Método 1: Descargar y extraer .tar.gz
	bpytop_tar_url := "https://github.com/aristocratos/bpytop/archive/refs/tags/v1.0.68.tar.gz"
	bpytop_cmd_tar := exec.Command("curl", "-L", bpytop_tar_url, "-o", "package.tar.gz")
	err := bpytop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bpytop_zip_url := "https://github.com/aristocratos/bpytop/archive/refs/tags/v1.0.68.zip"
	bpytop_cmd_zip := exec.Command("curl", "-L", bpytop_zip_url, "-o", "package.zip")
	err = bpytop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bpytop_bin_url := "https://github.com/aristocratos/bpytop/archive/refs/tags/v1.0.68.bin"
	bpytop_cmd_bin := exec.Command("curl", "-L", bpytop_bin_url, "-o", "binary.bin")
	err = bpytop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bpytop_src_url := "https://github.com/aristocratos/bpytop/archive/refs/tags/v1.0.68.src.tar.gz"
	bpytop_cmd_src := exec.Command("curl", "-L", bpytop_src_url, "-o", "source.tar.gz")
	err = bpytop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bpytop_cmd_direct := exec.Command("./binary")
	err = bpytop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: osx-cpu-temp")
	exec.Command("latte", "install", "osx-cpu-temp").Run()
}
