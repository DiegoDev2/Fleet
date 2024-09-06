package main

// Vimpc - Ncurses based mpd client with vi like key bindings
// Homepage: https://sourceforge.net/projects/vimpc/

import (
	"fmt"
	
	"os/exec"
)

func installVimpc() {
	// Método 1: Descargar y extraer .tar.gz
	vimpc_tar_url := "https://github.com/boysetsfrog/vimpc/archive/refs/tags/v0.09.2.tar.gz"
	vimpc_cmd_tar := exec.Command("curl", "-L", vimpc_tar_url, "-o", "package.tar.gz")
	err := vimpc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vimpc_zip_url := "https://github.com/boysetsfrog/vimpc/archive/refs/tags/v0.09.2.zip"
	vimpc_cmd_zip := exec.Command("curl", "-L", vimpc_zip_url, "-o", "package.zip")
	err = vimpc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vimpc_bin_url := "https://github.com/boysetsfrog/vimpc/archive/refs/tags/v0.09.2.bin"
	vimpc_cmd_bin := exec.Command("curl", "-L", vimpc_bin_url, "-o", "binary.bin")
	err = vimpc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vimpc_src_url := "https://github.com/boysetsfrog/vimpc/archive/refs/tags/v0.09.2.src.tar.gz"
	vimpc_cmd_src := exec.Command("curl", "-L", vimpc_src_url, "-o", "source.tar.gz")
	err = vimpc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vimpc_cmd_direct := exec.Command("./binary")
	err = vimpc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libmpdclient")
exec.Command("latte", "install", "libmpdclient")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: taglib")
exec.Command("latte", "install", "taglib")
}
