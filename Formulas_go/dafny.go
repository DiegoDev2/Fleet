package main

// Dafny - Verification-aware programming language
// Homepage: https://github.com/dafny-lang/dafny/blob/master/README.md

import (
	"fmt"
	
	"os/exec"
)

func installDafny() {
	// Método 1: Descargar y extraer .tar.gz
	dafny_tar_url := "https://github.com/dafny-lang/dafny/archive/refs/tags/v4.8.0.tar.gz"
	dafny_cmd_tar := exec.Command("curl", "-L", dafny_tar_url, "-o", "package.tar.gz")
	err := dafny_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dafny_zip_url := "https://github.com/dafny-lang/dafny/archive/refs/tags/v4.8.0.zip"
	dafny_cmd_zip := exec.Command("curl", "-L", dafny_zip_url, "-o", "package.zip")
	err = dafny_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dafny_bin_url := "https://github.com/dafny-lang/dafny/archive/refs/tags/v4.8.0.bin"
	dafny_cmd_bin := exec.Command("curl", "-L", dafny_bin_url, "-o", "binary.bin")
	err = dafny_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dafny_src_url := "https://github.com/dafny-lang/dafny/archive/refs/tags/v4.8.0.src.tar.gz"
	dafny_cmd_src := exec.Command("curl", "-L", dafny_src_url, "-o", "source.tar.gz")
	err = dafny_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dafny_cmd_direct := exec.Command("./binary")
	err = dafny_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dotnet@6")
exec.Command("latte", "install", "dotnet@6")
	fmt.Println("Instalando dependencia: openjdk@17")
exec.Command("latte", "install", "openjdk@17")
	fmt.Println("Instalando dependencia: z3")
exec.Command("latte", "install", "z3")
}
