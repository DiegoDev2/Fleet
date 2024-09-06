package main

// Libnids - Implements E-component of network intrusion detection system
// Homepage: https://libnids.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibnids() {
	// Método 1: Descargar y extraer .tar.gz
	libnids_tar_url := "https://downloads.sourceforge.net/project/libnids/libnids/1.24/libnids-1.24.tar.gz"
	libnids_cmd_tar := exec.Command("curl", "-L", libnids_tar_url, "-o", "package.tar.gz")
	err := libnids_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnids_zip_url := "https://downloads.sourceforge.net/project/libnids/libnids/1.24/libnids-1.24.zip"
	libnids_cmd_zip := exec.Command("curl", "-L", libnids_zip_url, "-o", "package.zip")
	err = libnids_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnids_bin_url := "https://downloads.sourceforge.net/project/libnids/libnids/1.24/libnids-1.24.bin"
	libnids_cmd_bin := exec.Command("curl", "-L", libnids_bin_url, "-o", "binary.bin")
	err = libnids_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnids_src_url := "https://downloads.sourceforge.net/project/libnids/libnids/1.24/libnids-1.24.src.tar.gz"
	libnids_cmd_src := exec.Command("curl", "-L", libnids_src_url, "-o", "source.tar.gz")
	err = libnids_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnids_cmd_direct := exec.Command("./binary")
	err = libnids_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libnet")
	exec.Command("latte", "install", "libnet").Run()
}
