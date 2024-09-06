package main

// SstpClient - SSTP (Microsoft's Remote Access Solution for PPP over SSL) client
// Homepage: https://sstp-client.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSstpClient() {
	// Método 1: Descargar y extraer .tar.gz
	sstpclient_tar_url := "https://gitlab.com/sstp-project/sstp-client/-/releases/1.0.20/downloads/dist-gzip/sstp-client-1.0.20.tar.gz"
	sstpclient_cmd_tar := exec.Command("curl", "-L", sstpclient_tar_url, "-o", "package.tar.gz")
	err := sstpclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sstpclient_zip_url := "https://gitlab.com/sstp-project/sstp-client/-/releases/1.0.20/downloads/dist-gzip/sstp-client-1.0.20.zip"
	sstpclient_cmd_zip := exec.Command("curl", "-L", sstpclient_zip_url, "-o", "package.zip")
	err = sstpclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sstpclient_bin_url := "https://gitlab.com/sstp-project/sstp-client/-/releases/1.0.20/downloads/dist-gzip/sstp-client-1.0.20.bin"
	sstpclient_cmd_bin := exec.Command("curl", "-L", sstpclient_bin_url, "-o", "binary.bin")
	err = sstpclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sstpclient_src_url := "https://gitlab.com/sstp-project/sstp-client/-/releases/1.0.20/downloads/dist-gzip/sstp-client-1.0.20.src.tar.gz"
	sstpclient_cmd_src := exec.Command("curl", "-L", sstpclient_src_url, "-o", "source.tar.gz")
	err = sstpclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sstpclient_cmd_direct := exec.Command("./binary")
	err = sstpclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
