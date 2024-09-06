package main

// NeovimQt - Neovim GUI, in Qt
// Homepage: https://github.com/equalsraf/neovim-qt

import (
	"fmt"
	
	"os/exec"
)

func installNeovimQt() {
	// Método 1: Descargar y extraer .tar.gz
	neovimqt_tar_url := "https://github.com/equalsraf/neovim-qt/archive/refs/tags/v0.2.18.tar.gz"
	neovimqt_cmd_tar := exec.Command("curl", "-L", neovimqt_tar_url, "-o", "package.tar.gz")
	err := neovimqt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neovimqt_zip_url := "https://github.com/equalsraf/neovim-qt/archive/refs/tags/v0.2.18.zip"
	neovimqt_cmd_zip := exec.Command("curl", "-L", neovimqt_zip_url, "-o", "package.zip")
	err = neovimqt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neovimqt_bin_url := "https://github.com/equalsraf/neovim-qt/archive/refs/tags/v0.2.18.bin"
	neovimqt_cmd_bin := exec.Command("curl", "-L", neovimqt_bin_url, "-o", "binary.bin")
	err = neovimqt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neovimqt_src_url := "https://github.com/equalsraf/neovim-qt/archive/refs/tags/v0.2.18.src.tar.gz"
	neovimqt_cmd_src := exec.Command("curl", "-L", neovimqt_src_url, "-o", "source.tar.gz")
	err = neovimqt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neovimqt_cmd_direct := exec.Command("./binary")
	err = neovimqt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: msgpack")
exec.Command("latte", "install", "msgpack")
	fmt.Println("Instalando dependencia: neovim")
exec.Command("latte", "install", "neovim")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
}
