package main

// Ncspot - Cross-platform ncurses Spotify client written in Rust
// Homepage: https://github.com/hrkfdn/ncspot

import (
	"fmt"
	
	"os/exec"
)

func installNcspot() {
	// Método 1: Descargar y extraer .tar.gz
	ncspot_tar_url := "https://github.com/hrkfdn/ncspot/archive/refs/tags/v1.1.2.tar.gz"
	ncspot_cmd_tar := exec.Command("curl", "-L", ncspot_tar_url, "-o", "package.tar.gz")
	err := ncspot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncspot_zip_url := "https://github.com/hrkfdn/ncspot/archive/refs/tags/v1.1.2.zip"
	ncspot_cmd_zip := exec.Command("curl", "-L", ncspot_zip_url, "-o", "package.zip")
	err = ncspot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncspot_bin_url := "https://github.com/hrkfdn/ncspot/archive/refs/tags/v1.1.2.bin"
	ncspot_cmd_bin := exec.Command("curl", "-L", ncspot_bin_url, "-o", "binary.bin")
	err = ncspot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncspot_src_url := "https://github.com/hrkfdn/ncspot/archive/refs/tags/v1.1.2.src.tar.gz"
	ncspot_cmd_src := exec.Command("curl", "-L", ncspot_src_url, "-o", "source.tar.gz")
	err = ncspot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncspot_cmd_direct := exec.Command("./binary")
	err = ncspot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: portaudio")
	exec.Command("latte", "install", "portaudio").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: dbus")
	exec.Command("latte", "install", "dbus").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
