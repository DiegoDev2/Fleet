package main

// LinuxPam - Pluggable Authentication Modules for Linux
// Homepage: https://github.com/linux-pam/linux-pam

import (
	"fmt"
	
	"os/exec"
)

func installLinuxPam() {
	// Método 1: Descargar y extraer .tar.gz
	linuxpam_tar_url := "https://github.com/linux-pam/linux-pam/releases/download/v1.6.1/Linux-PAM-1.6.1.tar.xz"
	linuxpam_cmd_tar := exec.Command("curl", "-L", linuxpam_tar_url, "-o", "package.tar.gz")
	err := linuxpam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linuxpam_zip_url := "https://github.com/linux-pam/linux-pam/releases/download/v1.6.1/Linux-PAM-1.6.1.tar.xz"
	linuxpam_cmd_zip := exec.Command("curl", "-L", linuxpam_zip_url, "-o", "package.zip")
	err = linuxpam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linuxpam_bin_url := "https://github.com/linux-pam/linux-pam/releases/download/v1.6.1/Linux-PAM-1.6.1.tar.xz"
	linuxpam_cmd_bin := exec.Command("curl", "-L", linuxpam_bin_url, "-o", "binary.bin")
	err = linuxpam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linuxpam_src_url := "https://github.com/linux-pam/linux-pam/releases/download/v1.6.1/Linux-PAM-1.6.1.tar.xz"
	linuxpam_cmd_src := exec.Command("curl", "-L", linuxpam_src_url, "-o", "source.tar.gz")
	err = linuxpam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linuxpam_cmd_direct := exec.Command("./binary")
	err = linuxpam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libnsl")
exec.Command("latte", "install", "libnsl")
	fmt.Println("Instalando dependencia: libprelude")
exec.Command("latte", "install", "libprelude")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
	fmt.Println("Instalando dependencia: libxcrypt")
exec.Command("latte", "install", "libxcrypt")
}
