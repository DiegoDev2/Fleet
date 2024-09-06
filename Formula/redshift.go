package main

// Redshift - Adjust color temperature of your screen according to your surroundings
// Homepage: http://jonls.dk/redshift/

import (
	"fmt"
	
	"os/exec"
)

func installRedshift() {
	// Método 1: Descargar y extraer .tar.gz
	redshift_tar_url := "https://github.com/jonls/redshift/releases/download/v1.12/redshift-1.12.tar.xz"
	redshift_cmd_tar := exec.Command("curl", "-L", redshift_tar_url, "-o", "package.tar.gz")
	err := redshift_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redshift_zip_url := "https://github.com/jonls/redshift/releases/download/v1.12/redshift-1.12.tar.xz"
	redshift_cmd_zip := exec.Command("curl", "-L", redshift_zip_url, "-o", "package.zip")
	err = redshift_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redshift_bin_url := "https://github.com/jonls/redshift/releases/download/v1.12/redshift-1.12.tar.xz"
	redshift_cmd_bin := exec.Command("curl", "-L", redshift_bin_url, "-o", "binary.bin")
	err = redshift_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redshift_src_url := "https://github.com/jonls/redshift/releases/download/v1.12/redshift-1.12.tar.xz"
	redshift_cmd_src := exec.Command("curl", "-L", redshift_src_url, "-o", "source.tar.gz")
	err = redshift_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redshift_cmd_direct := exec.Command("./binary")
	err = redshift_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
