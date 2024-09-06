package main

// Neovide - No Nonsense Neovim Client in Rust
// Homepage: https://github.com/neovide/neovide

import (
	"fmt"
	
	"os/exec"
)

func installNeovide() {
	// Método 1: Descargar y extraer .tar.gz
	neovide_tar_url := "https://github.com/neovide/neovide/archive/refs/tags/0.13.3.tar.gz"
	neovide_cmd_tar := exec.Command("curl", "-L", neovide_tar_url, "-o", "package.tar.gz")
	err := neovide_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neovide_zip_url := "https://github.com/neovide/neovide/archive/refs/tags/0.13.3.zip"
	neovide_cmd_zip := exec.Command("curl", "-L", neovide_zip_url, "-o", "package.zip")
	err = neovide_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neovide_bin_url := "https://github.com/neovide/neovide/archive/refs/tags/0.13.3.bin"
	neovide_cmd_bin := exec.Command("curl", "-L", neovide_bin_url, "-o", "binary.bin")
	err = neovide_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neovide_src_url := "https://github.com/neovide/neovide/archive/refs/tags/0.13.3.src.tar.gz"
	neovide_cmd_src := exec.Command("curl", "-L", neovide_src_url, "-o", "source.tar.gz")
	err = neovide_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neovide_cmd_direct := exec.Command("./binary")
	err = neovide_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: neovim")
	exec.Command("latte", "install", "neovim").Run()
	fmt.Println("Instalando dependencia: cargo-bundle")
	exec.Command("latte", "install", "cargo-bundle").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
}
