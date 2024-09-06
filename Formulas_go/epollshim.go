package main

// EpollShim - Small epoll implementation using kqueue
// Homepage: https://github.com/jiixyj/epoll-shim

import (
	"fmt"
	
	"os/exec"
)

func installEpollShim() {
	// Método 1: Descargar y extraer .tar.gz
	epollshim_tar_url := "https://github.com/jiixyj/epoll-shim/archive/refs/tags/v0.0.20240608.tar.gz"
	epollshim_cmd_tar := exec.Command("curl", "-L", epollshim_tar_url, "-o", "package.tar.gz")
	err := epollshim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	epollshim_zip_url := "https://github.com/jiixyj/epoll-shim/archive/refs/tags/v0.0.20240608.zip"
	epollshim_cmd_zip := exec.Command("curl", "-L", epollshim_zip_url, "-o", "package.zip")
	err = epollshim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	epollshim_bin_url := "https://github.com/jiixyj/epoll-shim/archive/refs/tags/v0.0.20240608.bin"
	epollshim_cmd_bin := exec.Command("curl", "-L", epollshim_bin_url, "-o", "binary.bin")
	err = epollshim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	epollshim_src_url := "https://github.com/jiixyj/epoll-shim/archive/refs/tags/v0.0.20240608.src.tar.gz"
	epollshim_cmd_src := exec.Command("curl", "-L", epollshim_src_url, "-o", "source.tar.gz")
	err = epollshim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	epollshim_cmd_direct := exec.Command("./binary")
	err = epollshim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
