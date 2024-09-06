package main

// Ser2net - Allow network connections to serial ports
// Homepage: https://ser2net.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSer2net() {
	// Método 1: Descargar y extraer .tar.gz
	ser2net_tar_url := "https://downloads.sourceforge.net/project/ser2net/ser2net/ser2net-4.6.2.tar.gz"
	ser2net_cmd_tar := exec.Command("curl", "-L", ser2net_tar_url, "-o", "package.tar.gz")
	err := ser2net_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ser2net_zip_url := "https://downloads.sourceforge.net/project/ser2net/ser2net/ser2net-4.6.2.zip"
	ser2net_cmd_zip := exec.Command("curl", "-L", ser2net_zip_url, "-o", "package.zip")
	err = ser2net_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ser2net_bin_url := "https://downloads.sourceforge.net/project/ser2net/ser2net/ser2net-4.6.2.bin"
	ser2net_cmd_bin := exec.Command("curl", "-L", ser2net_bin_url, "-o", "binary.bin")
	err = ser2net_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ser2net_src_url := "https://downloads.sourceforge.net/project/ser2net/ser2net/ser2net-4.6.2.src.tar.gz"
	ser2net_cmd_src := exec.Command("curl", "-L", ser2net_src_url, "-o", "source.tar.gz")
	err = ser2net_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ser2net_cmd_direct := exec.Command("./binary")
	err = ser2net_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gensio")
exec.Command("latte", "install", "gensio")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
