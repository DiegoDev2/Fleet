package main

// Nagios - Network monitoring and management system
// Homepage: https://www.nagios.org/

import (
	"fmt"
	
	"os/exec"
)

func installNagios() {
	// Método 1: Descargar y extraer .tar.gz
	nagios_tar_url := "https://downloads.sourceforge.net/project/nagios/nagios-4.x/nagios-4.5.4/nagios-4.5.4.tar.gz"
	nagios_cmd_tar := exec.Command("curl", "-L", nagios_tar_url, "-o", "package.tar.gz")
	err := nagios_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nagios_zip_url := "https://downloads.sourceforge.net/project/nagios/nagios-4.x/nagios-4.5.4/nagios-4.5.4.zip"
	nagios_cmd_zip := exec.Command("curl", "-L", nagios_zip_url, "-o", "package.zip")
	err = nagios_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nagios_bin_url := "https://downloads.sourceforge.net/project/nagios/nagios-4.x/nagios-4.5.4/nagios-4.5.4.bin"
	nagios_cmd_bin := exec.Command("curl", "-L", nagios_bin_url, "-o", "binary.bin")
	err = nagios_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nagios_src_url := "https://downloads.sourceforge.net/project/nagios/nagios-4.x/nagios-4.5.4/nagios-4.5.4.src.tar.gz"
	nagios_cmd_src := exec.Command("curl", "-L", nagios_src_url, "-o", "source.tar.gz")
	err = nagios_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nagios_cmd_direct := exec.Command("./binary")
	err = nagios_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gd")
	exec.Command("latte", "install", "gd").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
}
