package main

// Ntp - Network Time Protocol (NTP) Distribution
// Homepage: https://www.ntp.org

import (
	"fmt"
	
	"os/exec"
)

func installNtp() {
	// Método 1: Descargar y extraer .tar.gz
	ntp_tar_url := "https://downloads.nwtime.org/ntp/4.2.8/ntp-4.2.8p18.tar.gz"
	ntp_cmd_tar := exec.Command("curl", "-L", ntp_tar_url, "-o", "package.tar.gz")
	err := ntp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ntp_zip_url := "https://downloads.nwtime.org/ntp/4.2.8/ntp-4.2.8p18.zip"
	ntp_cmd_zip := exec.Command("curl", "-L", ntp_zip_url, "-o", "package.zip")
	err = ntp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ntp_bin_url := "https://downloads.nwtime.org/ntp/4.2.8/ntp-4.2.8p18.bin"
	ntp_cmd_bin := exec.Command("curl", "-L", ntp_bin_url, "-o", "binary.bin")
	err = ntp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ntp_src_url := "https://downloads.nwtime.org/ntp/4.2.8/ntp-4.2.8p18.src.tar.gz"
	ntp_cmd_src := exec.Command("curl", "-L", ntp_src_url, "-o", "source.tar.gz")
	err = ntp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ntp_cmd_direct := exec.Command("./binary")
	err = ntp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
