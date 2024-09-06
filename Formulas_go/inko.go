package main

// Inko - Safe and concurrent object-oriented programming language
// Homepage: https://inko-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installInko() {
	// Método 1: Descargar y extraer .tar.gz
	inko_tar_url := "https://releases.inko-lang.org/0.16.0.tar.gz"
	inko_cmd_tar := exec.Command("curl", "-L", inko_tar_url, "-o", "package.tar.gz")
	err := inko_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inko_zip_url := "https://releases.inko-lang.org/0.16.0.zip"
	inko_cmd_zip := exec.Command("curl", "-L", inko_zip_url, "-o", "package.zip")
	err = inko_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inko_bin_url := "https://releases.inko-lang.org/0.16.0.bin"
	inko_cmd_bin := exec.Command("curl", "-L", inko_bin_url, "-o", "binary.bin")
	err = inko_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inko_src_url := "https://releases.inko-lang.org/0.16.0.src.tar.gz"
	inko_cmd_src := exec.Command("curl", "-L", inko_src_url, "-o", "source.tar.gz")
	err = inko_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inko_cmd_direct := exec.Command("./binary")
	err = inko_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: z3")
exec.Command("latte", "install", "z3")
}
