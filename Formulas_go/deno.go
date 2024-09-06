package main

// Deno - Secure runtime for JavaScript and TypeScript
// Homepage: https://deno.com/

import (
	"fmt"
	
	"os/exec"
)

func installDeno() {
	// Método 1: Descargar y extraer .tar.gz
	deno_tar_url := "https://github.com/denoland/deno/releases/download/v1.45.5/deno_src.tar.gz"
	deno_cmd_tar := exec.Command("curl", "-L", deno_tar_url, "-o", "package.tar.gz")
	err := deno_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	deno_zip_url := "https://github.com/denoland/deno/releases/download/v1.45.5/deno_src.zip"
	deno_cmd_zip := exec.Command("curl", "-L", deno_zip_url, "-o", "package.zip")
	err = deno_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	deno_bin_url := "https://github.com/denoland/deno/releases/download/v1.45.5/deno_src.bin"
	deno_cmd_bin := exec.Command("curl", "-L", deno_bin_url, "-o", "binary.bin")
	err = deno_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	deno_src_url := "https://github.com/denoland/deno/releases/download/v1.45.5/deno_src.src.tar.gz"
	deno_cmd_src := exec.Command("curl", "-L", deno_src_url, "-o", "source.tar.gz")
	err = deno_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	deno_cmd_direct := exec.Command("./binary")
	err = deno_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
