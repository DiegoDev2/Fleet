package main

// Hyperkit - Toolkit for embedding hypervisor capabilities in your application
// Homepage: https://github.com/moby/hyperkit

import (
	"fmt"
	
	"os/exec"
)

func installHyperkit() {
	// Método 1: Descargar y extraer .tar.gz
	hyperkit_tar_url := "https://github.com/moby/hyperkit/archive/refs/tags/v0.20210107.tar.gz"
	hyperkit_cmd_tar := exec.Command("curl", "-L", hyperkit_tar_url, "-o", "package.tar.gz")
	err := hyperkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hyperkit_zip_url := "https://github.com/moby/hyperkit/archive/refs/tags/v0.20210107.zip"
	hyperkit_cmd_zip := exec.Command("curl", "-L", hyperkit_zip_url, "-o", "package.zip")
	err = hyperkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hyperkit_bin_url := "https://github.com/moby/hyperkit/archive/refs/tags/v0.20210107.bin"
	hyperkit_cmd_bin := exec.Command("curl", "-L", hyperkit_bin_url, "-o", "binary.bin")
	err = hyperkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hyperkit_src_url := "https://github.com/moby/hyperkit/archive/refs/tags/v0.20210107.src.tar.gz"
	hyperkit_cmd_src := exec.Command("curl", "-L", hyperkit_src_url, "-o", "source.tar.gz")
	err = hyperkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hyperkit_cmd_direct := exec.Command("./binary")
	err = hyperkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: opam")
exec.Command("latte", "install", "opam")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
}
