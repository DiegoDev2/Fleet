package main

// Lldpd - Implementation of IEEE 802.1ab (LLDP)
// Homepage: https://lldpd.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLldpd() {
	// Método 1: Descargar y extraer .tar.gz
	lldpd_tar_url := "https://media.luffy.cx/files/lldpd/lldpd-1.0.18.tar.gz"
	lldpd_cmd_tar := exec.Command("curl", "-L", lldpd_tar_url, "-o", "package.tar.gz")
	err := lldpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lldpd_zip_url := "https://media.luffy.cx/files/lldpd/lldpd-1.0.18.zip"
	lldpd_cmd_zip := exec.Command("curl", "-L", lldpd_zip_url, "-o", "package.zip")
	err = lldpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lldpd_bin_url := "https://media.luffy.cx/files/lldpd/lldpd-1.0.18.bin"
	lldpd_cmd_bin := exec.Command("curl", "-L", lldpd_bin_url, "-o", "binary.bin")
	err = lldpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lldpd_src_url := "https://media.luffy.cx/files/lldpd/lldpd-1.0.18.src.tar.gz"
	lldpd_cmd_src := exec.Command("curl", "-L", lldpd_src_url, "-o", "source.tar.gz")
	err = lldpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lldpd_cmd_direct := exec.Command("./binary")
	err = lldpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
