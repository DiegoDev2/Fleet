package main

// Nasm - Netwide Assembler (NASM) is an 80x86 assembler
// Homepage: https://www.nasm.us/

import (
	"fmt"
	
	"os/exec"
)

func installNasm() {
	// Método 1: Descargar y extraer .tar.gz
	nasm_tar_url := "https://www.nasm.us/pub/nasm/releasebuilds/2.16.03/nasm-2.16.03.tar.xz"
	nasm_cmd_tar := exec.Command("curl", "-L", nasm_tar_url, "-o", "package.tar.gz")
	err := nasm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nasm_zip_url := "https://www.nasm.us/pub/nasm/releasebuilds/2.16.03/nasm-2.16.03.tar.xz"
	nasm_cmd_zip := exec.Command("curl", "-L", nasm_zip_url, "-o", "package.zip")
	err = nasm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nasm_bin_url := "https://www.nasm.us/pub/nasm/releasebuilds/2.16.03/nasm-2.16.03.tar.xz"
	nasm_cmd_bin := exec.Command("curl", "-L", nasm_bin_url, "-o", "binary.bin")
	err = nasm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nasm_src_url := "https://www.nasm.us/pub/nasm/releasebuilds/2.16.03/nasm-2.16.03.tar.xz"
	nasm_cmd_src := exec.Command("curl", "-L", nasm_src_url, "-o", "source.tar.gz")
	err = nasm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nasm_cmd_direct := exec.Command("./binary")
	err = nasm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
}
