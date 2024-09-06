package main

// Megatools - Command-line client for Mega.co.nz
// Homepage: https://megatools.megous.com/

import (
	"fmt"
	
	"os/exec"
)

func installMegatools() {
	// Método 1: Descargar y extraer .tar.gz
	megatools_tar_url := "https://megatools.megous.com/builds/megatools-1.11.1.20230212.tar.gz"
	megatools_cmd_tar := exec.Command("curl", "-L", megatools_tar_url, "-o", "package.tar.gz")
	err := megatools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	megatools_zip_url := "https://megatools.megous.com/builds/megatools-1.11.1.20230212.zip"
	megatools_cmd_zip := exec.Command("curl", "-L", megatools_zip_url, "-o", "package.zip")
	err = megatools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	megatools_bin_url := "https://megatools.megous.com/builds/megatools-1.11.1.20230212.bin"
	megatools_cmd_bin := exec.Command("curl", "-L", megatools_bin_url, "-o", "binary.bin")
	err = megatools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	megatools_src_url := "https://megatools.megous.com/builds/megatools-1.11.1.20230212.src.tar.gz"
	megatools_cmd_src := exec.Command("curl", "-L", megatools_src_url, "-o", "source.tar.gz")
	err = megatools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	megatools_cmd_direct := exec.Command("./binary")
	err = megatools_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
