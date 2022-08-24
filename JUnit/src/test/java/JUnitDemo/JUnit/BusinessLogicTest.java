package JUnitDemo.JUnit;

import static org.junit.Assert.assertEquals;

import org.junit.Test;


public class BusinessLogicTest {
	
	
    
    @Test
    public void testBusinessLogic() {
    	boolean expect = true;
    	boolean actual = BusinessLogic.isPrime(3);
        assertEquals(expect,actual);
    }
    
    @Test
    public void testBusinessLogic2() {
    	boolean expect = false;
    	boolean actual = BusinessLogic.isPrime(4);
        assertEquals(expect,actual);
    }
    
    @Test
    public void testBusinessLogic3() {
    	boolean expect = false;
    	boolean actual = BusinessLogic.isPrime(-5);
        assertEquals(expect,actual);  
     }
}
