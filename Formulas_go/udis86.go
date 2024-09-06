package main

// Udis86 - Minimalistic disassembler library for x86
// Homepage: https://sourceforge.net/projects/udis86/

import (
	"fmt"
	
	"os/exec"
)

func installUdis86() {
	// Método 1: Descargar y extraer .tar.gz
	udis86_tar_url := "https://downloads.sourceforge.net/project/udis86/udis86/1.7/udis86-1.7.2.tar.gz"
	udis86_cmd_tar := exec.Command("curl", "-L", udis86_tar_url, "-o", "package.tar.gz")
	err := udis86_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	udis86_zip_url := "https://downloads.sourceforge.net/project/udis86/udis86/1.7/udis86-1.7.2.zip"
	udis86_cmd_zip := exec.Command("curl", "-L", udis86_zip_url, "-o", "package.zip")
	err = udis86_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	udis86_bin_url := "https://downloads.sourceforge.net/project/udis86/udis86/1.7/udis86-1.7.2.bin"
	udis86_cmd_bin := exec.Command("curl", "-L", udis86_bin_url, "-o", "binary.bin")
	err = udis86_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	udis86_src_url := "https://downloads.sourceforge.net/project/udis86/udis86/1.7/udis86-1.7.2.src.tar.gz"
	udis86_cmd_src := exec.Command("curl", "-L", udis86_src_url, "-o", "source.tar.gz")
	err = udis86_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	udis86_cmd_direct := exec.Command("./binary")
	err = udis86_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
