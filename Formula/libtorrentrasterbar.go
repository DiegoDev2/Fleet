package main

// LibtorrentRasterbar - C++ bittorrent library with Python bindings
// Homepage: https://www.libtorrent.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibtorrentRasterbar() {
	// Método 1: Descargar y extraer .tar.gz
	libtorrentrasterbar_tar_url := "https://github.com/arvidn/libtorrent/releases/download/v2.0.10/libtorrent-rasterbar-2.0.10.tar.gz"
	libtorrentrasterbar_cmd_tar := exec.Command("curl", "-L", libtorrentrasterbar_tar_url, "-o", "package.tar.gz")
	err := libtorrentrasterbar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtorrentrasterbar_zip_url := "https://github.com/arvidn/libtorrent/releases/download/v2.0.10/libtorrent-rasterbar-2.0.10.zip"
	libtorrentrasterbar_cmd_zip := exec.Command("curl", "-L", libtorrentrasterbar_zip_url, "-o", "package.zip")
	err = libtorrentrasterbar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtorrentrasterbar_bin_url := "https://github.com/arvidn/libtorrent/releases/download/v2.0.10/libtorrent-rasterbar-2.0.10.bin"
	libtorrentrasterbar_cmd_bin := exec.Command("curl", "-L", libtorrentrasterbar_bin_url, "-o", "binary.bin")
	err = libtorrentrasterbar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtorrentrasterbar_src_url := "https://github.com/arvidn/libtorrent/releases/download/v2.0.10/libtorrent-rasterbar-2.0.10.src.tar.gz"
	libtorrentrasterbar_cmd_src := exec.Command("curl", "-L", libtorrentrasterbar_src_url, "-o", "source.tar.gz")
	err = libtorrentrasterbar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtorrentrasterbar_cmd_direct := exec.Command("./binary")
	err = libtorrentrasterbar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: boost-python3")
	exec.Command("latte", "install", "boost-python3").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
