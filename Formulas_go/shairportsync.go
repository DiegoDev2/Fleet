package main

// ShairportSync - AirTunes emulator that adds multi-room capability
// Homepage: https://github.com/mikebrady/shairport-sync

import (
	"fmt"
	
	"os/exec"
)

func installShairportSync() {
	// Método 1: Descargar y extraer .tar.gz
	shairportsync_tar_url := "https://github.com/mikebrady/shairport-sync/archive/refs/tags/4.3.4.tar.gz"
	shairportsync_cmd_tar := exec.Command("curl", "-L", shairportsync_tar_url, "-o", "package.tar.gz")
	err := shairportsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shairportsync_zip_url := "https://github.com/mikebrady/shairport-sync/archive/refs/tags/4.3.4.zip"
	shairportsync_cmd_zip := exec.Command("curl", "-L", shairportsync_zip_url, "-o", "package.zip")
	err = shairportsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shairportsync_bin_url := "https://github.com/mikebrady/shairport-sync/archive/refs/tags/4.3.4.bin"
	shairportsync_cmd_bin := exec.Command("curl", "-L", shairportsync_bin_url, "-o", "binary.bin")
	err = shairportsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shairportsync_src_url := "https://github.com/mikebrady/shairport-sync/archive/refs/tags/4.3.4.src.tar.gz"
	shairportsync_cmd_src := exec.Command("curl", "-L", shairportsync_src_url, "-o", "source.tar.gz")
	err = shairportsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shairportsync_cmd_direct := exec.Command("./binary")
	err = shairportsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libao")
exec.Command("latte", "install", "libao")
	fmt.Println("Instalando dependencia: libconfig")
exec.Command("latte", "install", "libconfig")
	fmt.Println("Instalando dependencia: libdaemon")
exec.Command("latte", "install", "libdaemon")
	fmt.Println("Instalando dependencia: libsoxr")
exec.Command("latte", "install", "libsoxr")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
}
