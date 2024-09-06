package main

// PgpoolIi - PostgreSQL connection pool server
// Homepage: https://www.pgpool.net/mediawiki/index.php/Main_Page

import (
	"fmt"
	
	"os/exec"
)

func installPgpoolIi() {
	// Método 1: Descargar y extraer .tar.gz
	pgpoolii_tar_url := "https://www.pgpool.net/mediawiki/images/pgpool-II-4.5.3.tar.gz"
	pgpoolii_cmd_tar := exec.Command("curl", "-L", pgpoolii_tar_url, "-o", "package.tar.gz")
	err := pgpoolii_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgpoolii_zip_url := "https://www.pgpool.net/mediawiki/images/pgpool-II-4.5.3.zip"
	pgpoolii_cmd_zip := exec.Command("curl", "-L", pgpoolii_zip_url, "-o", "package.zip")
	err = pgpoolii_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgpoolii_bin_url := "https://www.pgpool.net/mediawiki/images/pgpool-II-4.5.3.bin"
	pgpoolii_cmd_bin := exec.Command("curl", "-L", pgpoolii_bin_url, "-o", "binary.bin")
	err = pgpoolii_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgpoolii_src_url := "https://www.pgpool.net/mediawiki/images/pgpool-II-4.5.3.src.tar.gz"
	pgpoolii_cmd_src := exec.Command("curl", "-L", pgpoolii_src_url, "-o", "source.tar.gz")
	err = pgpoolii_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgpoolii_cmd_direct := exec.Command("./binary")
	err = pgpoolii_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libmemcached")
	exec.Command("latte", "install", "libmemcached").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
}
