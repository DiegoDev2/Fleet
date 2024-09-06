package main

// Utimer - Multifunction timer tool
// Homepage: https://launchpad.net/utimer

import (
	"fmt"
	
	"os/exec"
)

func installUtimer() {
	// Método 1: Descargar y extraer .tar.gz
	utimer_tar_url := "https://launchpad.net/utimer/0.4/0.4/+download/utimer-0.4.tar.gz"
	utimer_cmd_tar := exec.Command("curl", "-L", utimer_tar_url, "-o", "package.tar.gz")
	err := utimer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	utimer_zip_url := "https://launchpad.net/utimer/0.4/0.4/+download/utimer-0.4.zip"
	utimer_cmd_zip := exec.Command("curl", "-L", utimer_zip_url, "-o", "package.zip")
	err = utimer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	utimer_bin_url := "https://launchpad.net/utimer/0.4/0.4/+download/utimer-0.4.bin"
	utimer_cmd_bin := exec.Command("curl", "-L", utimer_bin_url, "-o", "binary.bin")
	err = utimer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	utimer_src_url := "https://launchpad.net/utimer/0.4/0.4/+download/utimer-0.4.src.tar.gz"
	utimer_cmd_src := exec.Command("curl", "-L", utimer_src_url, "-o", "source.tar.gz")
	err = utimer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	utimer_cmd_direct := exec.Command("./binary")
	err = utimer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: perl-xml-parser")
	exec.Command("latte", "install", "perl-xml-parser").Run()
}
