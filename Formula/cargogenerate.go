package main

// CargoGenerate - Use pre-existing git repositories as templates
// Homepage: https://github.com/cargo-generate/cargo-generate

import (
	"fmt"
	
	"os/exec"
)

func installCargoGenerate() {
	// Método 1: Descargar y extraer .tar.gz
	cargogenerate_tar_url := "https://github.com/cargo-generate/cargo-generate/archive/refs/tags/v0.22.0.tar.gz"
	cargogenerate_cmd_tar := exec.Command("curl", "-L", cargogenerate_tar_url, "-o", "package.tar.gz")
	err := cargogenerate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargogenerate_zip_url := "https://github.com/cargo-generate/cargo-generate/archive/refs/tags/v0.22.0.zip"
	cargogenerate_cmd_zip := exec.Command("curl", "-L", cargogenerate_zip_url, "-o", "package.zip")
	err = cargogenerate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargogenerate_bin_url := "https://github.com/cargo-generate/cargo-generate/archive/refs/tags/v0.22.0.bin"
	cargogenerate_cmd_bin := exec.Command("curl", "-L", cargogenerate_bin_url, "-o", "binary.bin")
	err = cargogenerate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargogenerate_src_url := "https://github.com/cargo-generate/cargo-generate/archive/refs/tags/v0.22.0.src.tar.gz"
	cargogenerate_cmd_src := exec.Command("curl", "-L", cargogenerate_src_url, "-o", "source.tar.gz")
	err = cargogenerate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargogenerate_cmd_direct := exec.Command("./binary")
	err = cargogenerate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: libssh2")
	exec.Command("latte", "install", "libssh2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
