package main

// Ncmpcpp - Ncurses-based client for the Music Player Daemon
// Homepage: https://rybczak.net/ncmpcpp/

import (
	"fmt"
	
	"os/exec"
)

func installNcmpcpp() {
	// Método 1: Descargar y extraer .tar.gz
	ncmpcpp_tar_url := "https://rybczak.net/ncmpcpp/stable/ncmpcpp-0.9.2.tar.bz2"
	ncmpcpp_cmd_tar := exec.Command("curl", "-L", ncmpcpp_tar_url, "-o", "package.tar.gz")
	err := ncmpcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncmpcpp_zip_url := "https://rybczak.net/ncmpcpp/stable/ncmpcpp-0.9.2.tar.bz2"
	ncmpcpp_cmd_zip := exec.Command("curl", "-L", ncmpcpp_zip_url, "-o", "package.zip")
	err = ncmpcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncmpcpp_bin_url := "https://rybczak.net/ncmpcpp/stable/ncmpcpp-0.9.2.tar.bz2"
	ncmpcpp_cmd_bin := exec.Command("curl", "-L", ncmpcpp_bin_url, "-o", "binary.bin")
	err = ncmpcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncmpcpp_src_url := "https://rybczak.net/ncmpcpp/stable/ncmpcpp-0.9.2.tar.bz2"
	ncmpcpp_cmd_src := exec.Command("curl", "-L", ncmpcpp_src_url, "-o", "source.tar.gz")
	err = ncmpcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncmpcpp_cmd_direct := exec.Command("./binary")
	err = ncmpcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: libmpdclient")
	exec.Command("latte", "install", "libmpdclient").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: taglib")
	exec.Command("latte", "install", "taglib").Run()
}
