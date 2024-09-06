package main

// Sslh - Forward connections based on first data packet sent by client
// Homepage: https://www.rutschle.net/tech/sslh.shtml

import (
	"fmt"
	
	"os/exec"
)

func installSslh() {
	// Método 1: Descargar y extraer .tar.gz
	sslh_tar_url := "https://www.rutschle.net/tech/sslh/sslh-v2.1.2.tar.gz"
	sslh_cmd_tar := exec.Command("curl", "-L", sslh_tar_url, "-o", "package.tar.gz")
	err := sslh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sslh_zip_url := "https://www.rutschle.net/tech/sslh/sslh-v2.1.2.zip"
	sslh_cmd_zip := exec.Command("curl", "-L", sslh_zip_url, "-o", "package.zip")
	err = sslh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sslh_bin_url := "https://www.rutschle.net/tech/sslh/sslh-v2.1.2.bin"
	sslh_cmd_bin := exec.Command("curl", "-L", sslh_bin_url, "-o", "binary.bin")
	err = sslh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sslh_src_url := "https://www.rutschle.net/tech/sslh/sslh-v2.1.2.src.tar.gz"
	sslh_cmd_src := exec.Command("curl", "-L", sslh_src_url, "-o", "source.tar.gz")
	err = sslh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sslh_cmd_direct := exec.Command("./binary")
	err = sslh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libconfig")
	exec.Command("latte", "install", "libconfig").Run()
	fmt.Println("Instalando dependencia: libev")
	exec.Command("latte", "install", "libev").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
