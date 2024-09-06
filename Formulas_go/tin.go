package main

// Tin - Threaded, NNTP-, and spool-based UseNet newsreader
// Homepage: http://www.tin.org

import (
	"fmt"
	
	"os/exec"
)

func installTin() {
	// Método 1: Descargar y extraer .tar.gz
	tin_tar_url := "https://sunsite.icm.edu.pl/pub/unix/news/tin/v2.6/tin-2.6.3.tar.xz"
	tin_cmd_tar := exec.Command("curl", "-L", tin_tar_url, "-o", "package.tar.gz")
	err := tin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tin_zip_url := "https://sunsite.icm.edu.pl/pub/unix/news/tin/v2.6/tin-2.6.3.tar.xz"
	tin_cmd_zip := exec.Command("curl", "-L", tin_zip_url, "-o", "package.zip")
	err = tin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tin_bin_url := "https://sunsite.icm.edu.pl/pub/unix/news/tin/v2.6/tin-2.6.3.tar.xz"
	tin_cmd_bin := exec.Command("curl", "-L", tin_bin_url, "-o", "binary.bin")
	err = tin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tin_src_url := "https://sunsite.icm.edu.pl/pub/unix/news/tin/v2.6/tin-2.6.3.tar.xz"
	tin_cmd_src := exec.Command("curl", "-L", tin_src_url, "-o", "source.tar.gz")
	err = tin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tin_cmd_direct := exec.Command("./binary")
	err = tin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
