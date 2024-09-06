package main

// Ekhtml - Forgiving SAX-style HTML parser
// Homepage: https://ekhtml.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installEkhtml() {
	// Método 1: Descargar y extraer .tar.gz
	ekhtml_tar_url := "https://downloads.sourceforge.net/project/ekhtml/ekhtml/0.3.2/ekhtml-0.3.2.tar.gz"
	ekhtml_cmd_tar := exec.Command("curl", "-L", ekhtml_tar_url, "-o", "package.tar.gz")
	err := ekhtml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ekhtml_zip_url := "https://downloads.sourceforge.net/project/ekhtml/ekhtml/0.3.2/ekhtml-0.3.2.zip"
	ekhtml_cmd_zip := exec.Command("curl", "-L", ekhtml_zip_url, "-o", "package.zip")
	err = ekhtml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ekhtml_bin_url := "https://downloads.sourceforge.net/project/ekhtml/ekhtml/0.3.2/ekhtml-0.3.2.bin"
	ekhtml_cmd_bin := exec.Command("curl", "-L", ekhtml_bin_url, "-o", "binary.bin")
	err = ekhtml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ekhtml_src_url := "https://downloads.sourceforge.net/project/ekhtml/ekhtml/0.3.2/ekhtml-0.3.2.src.tar.gz"
	ekhtml_cmd_src := exec.Command("curl", "-L", ekhtml_src_url, "-o", "source.tar.gz")
	err = ekhtml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ekhtml_cmd_direct := exec.Command("./binary")
	err = ekhtml_cmd_direct.Run()
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
}
