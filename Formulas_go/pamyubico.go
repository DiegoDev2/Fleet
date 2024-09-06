package main

// PamYubico - Yubico pluggable authentication module
// Homepage: https://developers.yubico.com/yubico-pam/

import (
	"fmt"
	
	"os/exec"
)

func installPamYubico() {
	// Método 1: Descargar y extraer .tar.gz
	pamyubico_tar_url := "https://developers.yubico.com/yubico-pam/Releases/pam_yubico-2.27.tar.gz"
	pamyubico_cmd_tar := exec.Command("curl", "-L", pamyubico_tar_url, "-o", "package.tar.gz")
	err := pamyubico_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pamyubico_zip_url := "https://developers.yubico.com/yubico-pam/Releases/pam_yubico-2.27.zip"
	pamyubico_cmd_zip := exec.Command("curl", "-L", pamyubico_zip_url, "-o", "package.zip")
	err = pamyubico_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pamyubico_bin_url := "https://developers.yubico.com/yubico-pam/Releases/pam_yubico-2.27.bin"
	pamyubico_cmd_bin := exec.Command("curl", "-L", pamyubico_bin_url, "-o", "binary.bin")
	err = pamyubico_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pamyubico_src_url := "https://developers.yubico.com/yubico-pam/Releases/pam_yubico-2.27.src.tar.gz"
	pamyubico_cmd_src := exec.Command("curl", "-L", pamyubico_src_url, "-o", "source.tar.gz")
	err = pamyubico_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pamyubico_cmd_direct := exec.Command("./binary")
	err = pamyubico_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libyubikey")
exec.Command("latte", "install", "libyubikey")
	fmt.Println("Instalando dependencia: ykclient")
exec.Command("latte", "install", "ykclient")
	fmt.Println("Instalando dependencia: ykpers")
exec.Command("latte", "install", "ykpers")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
