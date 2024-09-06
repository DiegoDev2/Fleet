package main

// Bookloupe - List common formatting errors in a Project Gutenberg candidate file
// Homepage: http://www.juiblex.co.uk/pgdp/bookloupe/

import (
	"fmt"
	
	"os/exec"
)

func installBookloupe() {
	// Método 1: Descargar y extraer .tar.gz
	bookloupe_tar_url := "http://www.juiblex.co.uk/pgdp/bookloupe/bookloupe-2.0.tar.gz"
	bookloupe_cmd_tar := exec.Command("curl", "-L", bookloupe_tar_url, "-o", "package.tar.gz")
	err := bookloupe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bookloupe_zip_url := "http://www.juiblex.co.uk/pgdp/bookloupe/bookloupe-2.0.zip"
	bookloupe_cmd_zip := exec.Command("curl", "-L", bookloupe_zip_url, "-o", "package.zip")
	err = bookloupe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bookloupe_bin_url := "http://www.juiblex.co.uk/pgdp/bookloupe/bookloupe-2.0.bin"
	bookloupe_cmd_bin := exec.Command("curl", "-L", bookloupe_bin_url, "-o", "binary.bin")
	err = bookloupe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bookloupe_src_url := "http://www.juiblex.co.uk/pgdp/bookloupe/bookloupe-2.0.src.tar.gz"
	bookloupe_cmd_src := exec.Command("curl", "-L", bookloupe_src_url, "-o", "source.tar.gz")
	err = bookloupe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bookloupe_cmd_direct := exec.Command("./binary")
	err = bookloupe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
