package main

// AgePluginYubikey - Plugin for encrypting files with age and PIV tokens such as YubiKeys
// Homepage: https://github.com/str4d/age-plugin-yubikey

import (
	"fmt"
	
	"os/exec"
)

func installAgePluginYubikey() {
	// Método 1: Descargar y extraer .tar.gz
	agepluginyubikey_tar_url := "https://github.com/str4d/age-plugin-yubikey/archive/refs/tags/v0.5.0.tar.gz"
	agepluginyubikey_cmd_tar := exec.Command("curl", "-L", agepluginyubikey_tar_url, "-o", "package.tar.gz")
	err := agepluginyubikey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	agepluginyubikey_zip_url := "https://github.com/str4d/age-plugin-yubikey/archive/refs/tags/v0.5.0.zip"
	agepluginyubikey_cmd_zip := exec.Command("curl", "-L", agepluginyubikey_zip_url, "-o", "package.zip")
	err = agepluginyubikey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	agepluginyubikey_bin_url := "https://github.com/str4d/age-plugin-yubikey/archive/refs/tags/v0.5.0.bin"
	agepluginyubikey_cmd_bin := exec.Command("curl", "-L", agepluginyubikey_bin_url, "-o", "binary.bin")
	err = agepluginyubikey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	agepluginyubikey_src_url := "https://github.com/str4d/age-plugin-yubikey/archive/refs/tags/v0.5.0.src.tar.gz"
	agepluginyubikey_cmd_src := exec.Command("curl", "-L", agepluginyubikey_src_url, "-o", "source.tar.gz")
	err = agepluginyubikey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	agepluginyubikey_cmd_direct := exec.Command("./binary")
	err = agepluginyubikey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
