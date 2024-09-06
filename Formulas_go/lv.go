package main

// Lv - Powerful multi-lingual file viewer/grep
// Homepage: https://web.archive.org/web/20160310122517/www.ff.iij4u.or.jp/~nrt/lv/

import (
	"fmt"
	
	"os/exec"
)

func installLv() {
	// Método 1: Descargar y extraer .tar.gz
	lv_tar_url := "https://web.archive.org/web/20150915000000/www.ff.iij4u.or.jp/~nrt/freeware/lv451.tar.gz"
	lv_cmd_tar := exec.Command("curl", "-L", lv_tar_url, "-o", "package.tar.gz")
	err := lv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lv_zip_url := "https://web.archive.org/web/20150915000000/www.ff.iij4u.or.jp/~nrt/freeware/lv451.zip"
	lv_cmd_zip := exec.Command("curl", "-L", lv_zip_url, "-o", "package.zip")
	err = lv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lv_bin_url := "https://web.archive.org/web/20150915000000/www.ff.iij4u.or.jp/~nrt/freeware/lv451.bin"
	lv_cmd_bin := exec.Command("curl", "-L", lv_bin_url, "-o", "binary.bin")
	err = lv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lv_src_url := "https://web.archive.org/web/20150915000000/www.ff.iij4u.or.jp/~nrt/freeware/lv451.src.tar.gz"
	lv_cmd_src := exec.Command("curl", "-L", lv_src_url, "-o", "source.tar.gz")
	err = lv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lv_cmd_direct := exec.Command("./binary")
	err = lv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gzip")
exec.Command("latte", "install", "gzip")
}
