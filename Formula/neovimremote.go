package main

// NeovimRemote - Control nvim processes using `nvr` command-line tool
// Homepage: https://github.com/mhinz/neovim-remote

import (
	"fmt"
	
	"os/exec"
)

func installNeovimRemote() {
	// Método 1: Descargar y extraer .tar.gz
	neovimremote_tar_url := "https://files.pythonhosted.org/packages/69/50/4fe9ef6fd794929ceae73e476ac8a4ddbf3b0913fa248d834c9bb72978b7/neovim-remote-2.5.1.tar.gz"
	neovimremote_cmd_tar := exec.Command("curl", "-L", neovimremote_tar_url, "-o", "package.tar.gz")
	err := neovimremote_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neovimremote_zip_url := "https://files.pythonhosted.org/packages/69/50/4fe9ef6fd794929ceae73e476ac8a4ddbf3b0913fa248d834c9bb72978b7/neovim-remote-2.5.1.zip"
	neovimremote_cmd_zip := exec.Command("curl", "-L", neovimremote_zip_url, "-o", "package.zip")
	err = neovimremote_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neovimremote_bin_url := "https://files.pythonhosted.org/packages/69/50/4fe9ef6fd794929ceae73e476ac8a4ddbf3b0913fa248d834c9bb72978b7/neovim-remote-2.5.1.bin"
	neovimremote_cmd_bin := exec.Command("curl", "-L", neovimremote_bin_url, "-o", "binary.bin")
	err = neovimremote_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neovimremote_src_url := "https://files.pythonhosted.org/packages/69/50/4fe9ef6fd794929ceae73e476ac8a4ddbf3b0913fa248d834c9bb72978b7/neovim-remote-2.5.1.src.tar.gz"
	neovimremote_cmd_src := exec.Command("curl", "-L", neovimremote_src_url, "-o", "source.tar.gz")
	err = neovimremote_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neovimremote_cmd_direct := exec.Command("./binary")
	err = neovimremote_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: neovim")
	exec.Command("latte", "install", "neovim").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
