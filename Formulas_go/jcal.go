package main

// Jcal - UNIX-cal-like tool to display Jalali calendar
// Homepage: https://savannah.nongnu.org/projects/jcal/

import (
	"fmt"
	
	"os/exec"
)

func installJcal() {
	// Método 1: Descargar y extraer .tar.gz
	jcal_tar_url := "https://download.savannah.gnu.org/releases/jcal/jcal-0.4.1.tar.gz"
	jcal_cmd_tar := exec.Command("curl", "-L", jcal_tar_url, "-o", "package.tar.gz")
	err := jcal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jcal_zip_url := "https://download.savannah.gnu.org/releases/jcal/jcal-0.4.1.zip"
	jcal_cmd_zip := exec.Command("curl", "-L", jcal_zip_url, "-o", "package.zip")
	err = jcal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jcal_bin_url := "https://download.savannah.gnu.org/releases/jcal/jcal-0.4.1.bin"
	jcal_cmd_bin := exec.Command("curl", "-L", jcal_bin_url, "-o", "binary.bin")
	err = jcal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jcal_src_url := "https://download.savannah.gnu.org/releases/jcal/jcal-0.4.1.src.tar.gz"
	jcal_cmd_src := exec.Command("curl", "-L", jcal_src_url, "-o", "source.tar.gz")
	err = jcal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jcal_cmd_direct := exec.Command("./binary")
	err = jcal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
