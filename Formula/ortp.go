package main

// Ortp - Real-time transport protocol (RTP, RFC3550) library
// Homepage: https://linphone.org/

import (
	"fmt"
	
	"os/exec"
)

func installOrtp() {
	// Método 1: Descargar y extraer .tar.gz
	ortp_tar_url := "https://gitlab.linphone.org/BC/public/ortp/-/archive/5.3.79/ortp-5.3.79.tar.bz2"
	ortp_cmd_tar := exec.Command("curl", "-L", ortp_tar_url, "-o", "package.tar.gz")
	err := ortp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ortp_zip_url := "https://gitlab.linphone.org/BC/public/ortp/-/archive/5.3.79/ortp-5.3.79.tar.bz2"
	ortp_cmd_zip := exec.Command("curl", "-L", ortp_zip_url, "-o", "package.zip")
	err = ortp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ortp_bin_url := "https://gitlab.linphone.org/BC/public/ortp/-/archive/5.3.79/ortp-5.3.79.tar.bz2"
	ortp_cmd_bin := exec.Command("curl", "-L", ortp_bin_url, "-o", "binary.bin")
	err = ortp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ortp_src_url := "https://gitlab.linphone.org/BC/public/ortp/-/archive/5.3.79/ortp-5.3.79.tar.bz2"
	ortp_cmd_src := exec.Command("curl", "-L", ortp_src_url, "-o", "source.tar.gz")
	err = ortp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ortp_cmd_direct := exec.Command("./binary")
	err = ortp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: mbedtls")
	exec.Command("latte", "install", "mbedtls").Run()
}
