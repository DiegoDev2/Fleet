package main

// Ganglia - Scalable distributed monitoring system
// Homepage: https://ganglia.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGanglia() {
	// Método 1: Descargar y extraer .tar.gz
	ganglia_tar_url := "https://downloads.sourceforge.net/project/ganglia/ganglia%20monitoring%20core/3.7.2/ganglia-3.7.2.tar.gz"
	ganglia_cmd_tar := exec.Command("curl", "-L", ganglia_tar_url, "-o", "package.tar.gz")
	err := ganglia_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ganglia_zip_url := "https://downloads.sourceforge.net/project/ganglia/ganglia%20monitoring%20core/3.7.2/ganglia-3.7.2.zip"
	ganglia_cmd_zip := exec.Command("curl", "-L", ganglia_zip_url, "-o", "package.zip")
	err = ganglia_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ganglia_bin_url := "https://downloads.sourceforge.net/project/ganglia/ganglia%20monitoring%20core/3.7.2/ganglia-3.7.2.bin"
	ganglia_cmd_bin := exec.Command("curl", "-L", ganglia_bin_url, "-o", "binary.bin")
	err = ganglia_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ganglia_src_url := "https://downloads.sourceforge.net/project/ganglia/ganglia%20monitoring%20core/3.7.2/ganglia-3.7.2.src.tar.gz"
	ganglia_cmd_src := exec.Command("curl", "-L", ganglia_src_url, "-o", "source.tar.gz")
	err = ganglia_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ganglia_cmd_direct := exec.Command("./binary")
	err = ganglia_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: apr")
	exec.Command("latte", "install", "apr").Run()
	fmt.Println("Instalando dependencia: confuse")
	exec.Command("latte", "install", "confuse").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: rrdtool")
	exec.Command("latte", "install", "rrdtool").Run()
}
