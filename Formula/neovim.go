package main

// Neovim - Ambitious Vim-fork focused on extensibility and agility
// Homepage: https://neovim.io/

import (
	"fmt"
	
	"os/exec"
)

func installNeovim() {
	// Método 1: Descargar y extraer .tar.gz
	neovim_tar_url := "https://github.com/neovim/neovim/archive/refs/tags/v0.10.1.tar.gz"
	neovim_cmd_tar := exec.Command("curl", "-L", neovim_tar_url, "-o", "package.tar.gz")
	err := neovim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neovim_zip_url := "https://github.com/neovim/neovim/archive/refs/tags/v0.10.1.zip"
	neovim_cmd_zip := exec.Command("curl", "-L", neovim_zip_url, "-o", "package.zip")
	err = neovim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neovim_bin_url := "https://github.com/neovim/neovim/archive/refs/tags/v0.10.1.bin"
	neovim_cmd_bin := exec.Command("curl", "-L", neovim_bin_url, "-o", "binary.bin")
	err = neovim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neovim_src_url := "https://github.com/neovim/neovim/archive/refs/tags/v0.10.1.src.tar.gz"
	neovim_cmd_src := exec.Command("curl", "-L", neovim_src_url, "-o", "source.tar.gz")
	err = neovim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neovim_cmd_direct := exec.Command("./binary")
	err = neovim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libvterm")
	exec.Command("latte", "install", "libvterm").Run()
	fmt.Println("Instalando dependencia: msgpack")
	exec.Command("latte", "install", "msgpack").Run()
	fmt.Println("Instalando dependencia: utf8proc")
	exec.Command("latte", "install", "utf8proc").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: lpeg")
	exec.Command("latte", "install", "lpeg").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: luv")
	exec.Command("latte", "install", "luv").Run()
	fmt.Println("Instalando dependencia: tree-sitter")
	exec.Command("latte", "install", "tree-sitter").Run()
	fmt.Println("Instalando dependencia: unibilium")
	exec.Command("latte", "install", "unibilium").Run()
	fmt.Println("Instalando dependencia: libnsl")
	exec.Command("latte", "install", "libnsl").Run()
}
