package main

// Asm6809 - Cross assembler targeting the Motorola 6809 and Hitachi 6309
// Homepage: https://www.6809.org.uk/asm6809/

import (
	"fmt"
	
	"os/exec"
)

func installAsm6809() {
	// Método 1: Descargar y extraer .tar.gz
	asm6809_tar_url := "https://www.6809.org.uk/asm6809/dl/asm6809-2.13.tar.gz"
	asm6809_cmd_tar := exec.Command("curl", "-L", asm6809_tar_url, "-o", "package.tar.gz")
	err := asm6809_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asm6809_zip_url := "https://www.6809.org.uk/asm6809/dl/asm6809-2.13.zip"
	asm6809_cmd_zip := exec.Command("curl", "-L", asm6809_zip_url, "-o", "package.zip")
	err = asm6809_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asm6809_bin_url := "https://www.6809.org.uk/asm6809/dl/asm6809-2.13.bin"
	asm6809_cmd_bin := exec.Command("curl", "-L", asm6809_bin_url, "-o", "binary.bin")
	err = asm6809_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asm6809_src_url := "https://www.6809.org.uk/asm6809/dl/asm6809-2.13.src.tar.gz"
	asm6809_cmd_src := exec.Command("curl", "-L", asm6809_src_url, "-o", "source.tar.gz")
	err = asm6809_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asm6809_cmd_direct := exec.Command("./binary")
	err = asm6809_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
