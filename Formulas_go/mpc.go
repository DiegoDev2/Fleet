package main

// Mpc - Command-line music player client for mpd
// Homepage: https://www.musicpd.org/clients/mpc/

import (
	"fmt"
	
	"os/exec"
)

func installMpc() {
	// Método 1: Descargar y extraer .tar.gz
	mpc_tar_url := "https://www.musicpd.org/download/mpc/0/mpc-0.35.tar.xz"
	mpc_cmd_tar := exec.Command("curl", "-L", mpc_tar_url, "-o", "package.tar.gz")
	err := mpc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpc_zip_url := "https://www.musicpd.org/download/mpc/0/mpc-0.35.tar.xz"
	mpc_cmd_zip := exec.Command("curl", "-L", mpc_zip_url, "-o", "package.zip")
	err = mpc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpc_bin_url := "https://www.musicpd.org/download/mpc/0/mpc-0.35.tar.xz"
	mpc_cmd_bin := exec.Command("curl", "-L", mpc_bin_url, "-o", "binary.bin")
	err = mpc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpc_src_url := "https://www.musicpd.org/download/mpc/0/mpc-0.35.tar.xz"
	mpc_cmd_src := exec.Command("curl", "-L", mpc_src_url, "-o", "source.tar.gz")
	err = mpc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpc_cmd_direct := exec.Command("./binary")
	err = mpc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libmpdclient")
exec.Command("latte", "install", "libmpdclient")
}
